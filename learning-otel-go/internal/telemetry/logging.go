package telemetry

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Logger é uma estrutura personalizada para logs que integra com OpenTelemetry
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// NewLogger cria uma nova instância de Logger
func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info registra uma mensagem de informação
func (l *Logger) Info(ctx context.Context, msg string, keyValues ...interface{}) {
	// Adiciona informações de trace se disponíveis
	addTraceInfoToLog(ctx, l.infoLogger, msg, keyValues...)
}

// Error registra uma mensagem de erro
func (l *Logger) Error(ctx context.Context, msg string, keyValues ...interface{}) {
	// Marca o span atual com o erro (se existir)
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		span.SetStatus(codes.Error, msg)

		// Adiciona atributos do erro ao span
		for i := 0; i < len(keyValues); i += 2 {
			if i+1 < len(keyValues) {
				key, ok := keyValues[i].(string)
				if ok {
					span.SetAttributes(attribute.String(key, toString(keyValues[i+1])))
				}
			}
		}
	}

	// Registra no log
	addTraceInfoToLog(ctx, l.errorLogger, msg, keyValues...)
}

// Debug registra uma mensagem de depuração
func (l *Logger) Debug(ctx context.Context, msg string, keyValues ...interface{}) {
	// Adiciona informações de trace se disponíveis
	addTraceInfoToLog(ctx, l.debugLogger, msg, keyValues...)
}

// addTraceInfoToLog adiciona informações de trace ao log
func addTraceInfoToLog(ctx context.Context, logger *log.Logger, msg string, keyValues ...interface{}) {
	// Extrai o TraceID e SpanID do contexto, se estiver disponível
	span := trace.SpanFromContext(ctx)

	if span.IsRecording() {
		spanContext := span.SpanContext()
		if spanContext.IsValid() {
			// Adiciona os IDs de trace e span à mensagem
			msg = msg + " [trace_id=" + spanContext.TraceID().String() +
				" span_id=" + spanContext.SpanID().String() + "]"
		}
	}

	// Formata os pares chave-valor adicionais
	if len(keyValues) > 0 {
		pairs := make([]interface{}, 0, len(keyValues)+1)
		pairs = append(pairs, msg)
		pairs = append(pairs, keyValues...)
		logger.Println(pairs...)
	} else {
		logger.Println(msg)
	}
}

// toString converte um valor para string
func toString(value interface{}) string {
	if value == nil {
		return "<nil>"
	}
	if s, ok := value.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", value)
}
