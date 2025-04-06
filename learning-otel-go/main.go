package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"todo-api/internal/config"
	"todo-api/internal/core/task"
	"todo-api/internal/router"
	"todo-api/internal/telemetry"
)

func main() {
	// Carregar configurações
	var cfg = config.LoadConfig()

	// Inicializar OpenTelemetry
	collectorURL := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if collectorURL == "" {
		collectorURL = "localhost:4317" // Valor padrão para o coletor OTLP
	}

	cleanup, err := telemetry.InitTracer("todo-api", collectorURL)
	if err != nil {
		log.Printf("Erro ao inicializar OpenTelemetry: %v. Continuando sem telemetria.", err)
	} else {
		// Garantir que o tracer seja encerrado corretamente na saída
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := cleanup(ctx); err != nil {
				log.Printf("Erro ao limpar recursos de telemetria: %v", err)
			}
		}()
	}

	// Conectar ao banco de dados
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Banco indisponível:", err)
	}

	// Inicializar componentes
	repo := task.NewPgTaskRepository(db)
	service := task.NewTaskService(repo)
	r := router.NewRouter(service)

	// Adicionar middleware de telemetria
	handler := telemetry.TracingMiddleware(r)

	// Configurar e iniciar o servidor
	port := cfg.Server.Port
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Iniciar o servidor em uma goroutine
	go func() {
		log.Printf("Servidor rodando em http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Desligando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro durante shutdown do servidor: %v", err)
	}

	log.Println("Servidor desligado com sucesso")
}
