package telemetry

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// TracingMiddleware instrumenta requisições HTTP com OpenTelemetry
func TracingMiddleware(next http.Handler) http.Handler {
	tracer := otel.Tracer("http-middleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		// Wrapper do ResponseWriter para capturar o código de status
		wrapper := &responseWriter{w: w, status: http.StatusOK}

		// Passa o contexto de trace para o próximo handler
		next.ServeHTTP(wrapper, r.WithContext(ctx))

		// Adiciona informações de resposta ao span
		span.SetAttributes(
			attribute.Int("http.status_code", wrapper.status),
		)

		// Se o status é de erro, marca o span como erro
		if wrapper.status >= 400 {
			span.SetStatus(codes.Error, http.StatusText(wrapper.status))
		} else {
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
