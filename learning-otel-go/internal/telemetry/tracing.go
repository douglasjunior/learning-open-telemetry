package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitTracer inicializa o rastreamento do OpenTelemetry
func InitTracer(serviceName string, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	// Cria uma conexão com o coletor OTLP
	conn, err := grpc.DialContext(ctx, collectorURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com o coletor OTLP: %w", err)
	}

	// Cria o exportador de traces
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("falha ao criar exportador OTLP: %w", err)
	}

	// Cria um recurso que identifica a aplicação
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar recurso: %w", err)
	}

	// Cria um provedor de tracer
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	// Define o provedor global de tracer
	otel.SetTracerProvider(tracerProvider)

	// Define o propagador global
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Função para limpeza ao encerrar a aplicação
	cleanup := func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := tracerProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("falha ao desligar provedor de tracer: %w", err)
		}
		return nil
	}

	return cleanup, nil
}
