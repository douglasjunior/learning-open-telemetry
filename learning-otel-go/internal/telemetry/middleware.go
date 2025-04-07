package telemetry

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// TracingMiddleware instrumenta requisições HTTP com OpenTelemetry
func TracingMiddleware(next http.Handler) http.Handler {
	tracer := otel.Tracer("http-middleware")
	logger := NewLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Marca o início para cálculo de duração
		startTime := time.Now()

		// Extrai contexto de trace da requisição
		propagator := otel.GetTextMapPropagator()
		ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))

		// Inicia um novo span para esta requisição
		path := r.URL.Path
		method := r.Method
		ctx, span := tracer.Start(
			ctx,
			method+" "+path,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				attribute.String("http.method", method),
				attribute.String("http.url", path),
				attribute.String("http.user_agent", r.UserAgent()),
				attribute.String("http.remote_addr", r.RemoteAddr),
			),
		)
		defer span.End()

		// Log da requisição recebida
		logger.Info(ctx, "Requisição recebida",
			"method", method,
			"path", path,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
		)

		// Wrapper do ResponseWriter para capturar o código de status
		wrapper := &responseWriter{w: w, status: http.StatusOK}

		// Passa o contexto de trace para o próximo handler
		next.ServeHTTP(wrapper, r.WithContext(ctx))

		// Calcula a duração
		duration := float64(time.Since(startTime).Milliseconds())

		// Adiciona informações de resposta ao span
		span.SetAttributes(
			attribute.Int("http.status_code", wrapper.status),
			attribute.Float64("http.duration_ms", duration),
		)

		// Registra métricas (comentado até os pacotes de métricas serem adicionados)
		/*
			RecordHTTPRequest(ctx, method, path, wrapper.status, duration)
		*/

		// Log da resposta
		if wrapper.status >= 400 {
			logger.Error(ctx, "Erro na resposta HTTP",
				"method", method,
				"path", path,
				"status", wrapper.status,
				"duration_ms", duration,
			)
			span.SetStatus(codes.Error, http.StatusText(wrapper.status))
		} else {
			logger.Info(ctx, "Resposta enviada com sucesso",
				"method", method,
				"path", path,
				"status", wrapper.status,
				"duration_ms", duration,
			)
			span.SetStatus(codes.Ok, "")
		}
	})
}

// responseWriter é um wrapper de http.ResponseWriter que captura o código de status
type responseWriter struct {
	w      http.ResponseWriter
	status int
}

func (rw *responseWriter) Header() http.Header {
	return rw.w.Header()
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.w.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.w.WriteHeader(statusCode)
}
