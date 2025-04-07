package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"todo-api/internal/config"
	"todo-api/internal/core/task"
	"todo-api/internal/router"
	"todo-api/internal/telemetry"
)

// Application encapsula toda a lógica da aplicação
type Application struct {
	config *config.Config
	logger *telemetry.Logger
	server *http.Server
	db     *sql.DB
}

// New cria uma nova instância da aplicação
func New() *Application {
	return &Application{
		config: config.LoadConfig(),
		logger: telemetry.NewLogger(),
	}
}

// Initialize inicializa todos os componentes da aplicação
func (app *Application) Initialize(ctx context.Context) error {
	if err := app.setupTelemetry(ctx); err != nil {
		return fmt.Errorf("falha ao configurar telemetria: %w", err)
	}

	if err := app.setupDatabase(ctx); err != nil {
		return fmt.Errorf("falha ao configurar banco de dados: %w", err)
	}

	app.setupServer(ctx)
	return nil
}

// setupTelemetry inicializa os componentes de telemetria
func (app *Application) setupTelemetry(ctx context.Context) error {
	collectorURL := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	fmt.Println("collectorURL", collectorURL)
	if collectorURL == "" {
		collectorURL = "localhost:4317"
	}

	// Inicializar tracer
	traceCleanup, err := telemetry.InitTracer("todo-api", collectorURL)
	if err != nil {
		app.logger.Error(ctx, "Erro ao inicializar OpenTelemetry Tracer", "error", err)
		return err
	}

	// Registrar função de limpeza
	go func() {
		<-ctx.Done()
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := traceCleanup(cleanupCtx); err != nil {
			app.logger.Error(cleanupCtx, "Erro ao limpar recursos de tracer", "error", err)
		}
	}()

	// Inicializar métricas
	meterCleanup, err := telemetry.InitMeter("todo-api", collectorURL)
	if err != nil {
		app.logger.Error(ctx, "Erro ao inicializar OpenTelemetry Metrics", "error", err)
		return err
	}

	// Registrar função de limpeza
	go func() {
		<-ctx.Done()
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := meterCleanup(cleanupCtx); err != nil {
			app.logger.Error(cleanupCtx, "Erro ao limpar recursos de métricas", "error", err)
		}
	}()

	return nil
}

// setupDatabase configura a conexão com o banco de dados
func (app *Application) setupDatabase(ctx context.Context) error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		app.config.Database.User,
		app.config.Database.Password,
		app.config.Database.Host,
		app.config.Database.Port,
		app.config.Database.Name)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		app.logger.Error(ctx, "Erro ao conectar no banco", "error", err)
		return err
	}

	if err := db.Ping(); err != nil {
		app.logger.Error(ctx, "Banco indisponível", "error", err)
		return err
	}

	app.db = db
	app.logger.Info(ctx, "Conexão com o banco de dados estabelecida com sucesso")
	return nil
}

// setupServer configura o servidor HTTP
func (app *Application) setupServer(ctx context.Context) {
	repo := task.NewPgTaskRepository(app.db)
	service := task.NewTaskService(repo)
	r := router.NewRouter(service)

	// Adicionar middleware de telemetria
	handler := telemetry.TracingMiddleware(r)

	port := app.config.Server.Port
	app.server = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

// Start inicia o servidor e configura o graceful shutdown
func (app *Application) Start(ctx context.Context) error {
	// Contexto para gerenciar o ciclo de vida da aplicação
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Canal para receber sinais de término
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar o servidor em uma goroutine
	errChan := make(chan error, 1)
	go func() {
		port := app.config.Server.Port
		app.logger.Info(ctx, "Servidor iniciado", "port", port)
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Error(ctx, "Erro ao iniciar servidor", "error", err)
			errChan <- err
		}
	}()

	// Aguardar sinal de término ou erro
	select {
	case <-quit:
		app.logger.Info(ctx, "Desligando servidor...")
	case err := <-errChan:
		app.logger.Error(ctx, "Erro fatal no servidor", "error", err)
		return err
	}

	return app.Shutdown(ctx)
}

// Shutdown desliga o servidor de forma graciosa
func (app *Application) Shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := app.server.Shutdown(shutdownCtx); err != nil {
		app.logger.Error(shutdownCtx, "Erro durante shutdown do servidor", "error", err)
		return err
	}

	if app.db != nil {
		if err := app.db.Close(); err != nil {
			app.logger.Error(shutdownCtx, "Erro ao fechar conexão com banco de dados", "error", err)
			return err
		}
	}

	app.logger.Info(shutdownCtx, "Servidor desligado com sucesso")
	return nil
}
