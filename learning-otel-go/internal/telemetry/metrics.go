package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/resource"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// Variáveis globais de métricas para uso em todo o aplicativo
var (
	// Contador para requisições HTTP
	httpRequestCounter metric.Int64Counter
	// Histograma para duração de requisições
	httpRequestDuration metric.Float64Histogram
	// Contador para operações de banco de dados
	dbOperationCounter metric.Int64Counter
	// Histograma para duração de operações de banco de dados
	dbOperationDuration metric.Float64Histogram
)

// InitMeter inicializa o provedor de métricas do OpenTelemetry
func InitMeter(serviceName string, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	// Cria um recurso que identifica a aplicação
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar recurso: %w", err)
	}

	// Cria o exportador de métricas usando gRPC com o mesmo coletor dos traces
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(collectorURL),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar exportador de métricas: %w", err)
	}

	// Cria o leitor periódico de métricas que envia dados a cada 15 segundos
	reader := sdkmetric.NewPeriodicReader(exporter,
		sdkmetric.WithInterval(15*time.Second),
	)

	// Cria um provedor de métricas
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(reader),
	)

	// Define o provedor global de métricas
	otel.SetMeterProvider(meterProvider)

	// Cria um medidor de métricas específico para nossa aplicação
	meter := meterProvider.Meter(
		"todo-api",
		metric.WithInstrumentationVersion("0.1.0"),
	)

	// Inicializa as métricas que serão usadas pelo aplicativo
	initMetrics(meter)

	// Função para limpeza ao encerrar a aplicação
	cleanup := func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := meterProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("falha ao desligar provedor de métricas: %w", err)
		}
		return nil
	}

	return cleanup, nil
}

// Inicializa as métricas que serão usadas pelo aplicativo
func initMetrics(meter metric.Meter) {
	var err error

	// Contador de requisições HTTP
	httpRequestCounter, err = meter.Int64Counter(
		"http.requests.total",
		metric.WithDescription("Número total de requisições HTTP"),
		metric.WithUnit("1"),
	)
	if err != nil {
		otel.Handle(err)
	}

	// Histograma de duração de requisições HTTP
	httpRequestDuration, err = meter.Float64Histogram(
		"http.request.duration",
		metric.WithDescription("Duração das requisições HTTP"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		otel.Handle(err)
	}

	// Contador de operações de banco de dados
	dbOperationCounter, err = meter.Int64Counter(
		"db.operations.total",
		metric.WithDescription("Número total de operações de banco de dados"),
		metric.WithUnit("1"),
	)
	if err != nil {
		otel.Handle(err)
	}

	// Histograma de duração de operações de banco de dados
	dbOperationDuration, err = meter.Float64Histogram(
		"db.operation.duration",
		metric.WithDescription("Duração das operações de banco de dados"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		otel.Handle(err)
	}
}

// RecordHTTPRequest registra uma requisição HTTP
func RecordHTTPRequest(ctx context.Context, method, route string, statusCode int, duration float64) {
	httpRequestCounter.Add(ctx, 1,
		metric.WithAttributes(
			semconv.HTTPMethodKey.String(method),
			semconv.HTTPRouteKey.String(route),
			semconv.HTTPStatusCodeKey.Int(statusCode),
		),
	)

	httpRequestDuration.Record(ctx, duration,
		metric.WithAttributes(
			semconv.HTTPMethodKey.String(method),
			semconv.HTTPRouteKey.String(route),
			semconv.HTTPStatusCodeKey.Int(statusCode),
		),
	)
}

// RecordDBOperation registra uma operação de banco de dados
func RecordDBOperation(ctx context.Context, operation string, success bool, duration float64) {
	dbOperationCounter.Add(ctx, 1,
		metric.WithAttributes(
			semconv.DBOperationKey.String(operation),
			semconv.DBSystemKey.String("postgresql"),
			semconv.DBStatementKey.String(operation),
		),
	)

	dbOperationDuration.Record(ctx, duration,
		metric.WithAttributes(
			semconv.DBOperationKey.String(operation),
			semconv.DBSystemKey.String("postgresql"),
			semconv.DBStatementKey.String(operation),
		),
	)
}
