# API de Tarefas com OpenTelemetry em Go

Este projeto é um exemplo prático de como implementar telemetria usando OpenTelemetry em uma aplicação Go. A aplicação consiste em uma API de gerenciamento de tarefas (todo list) com uma arquitetura em camadas e persistência em PostgreSQL.

## Estrutura do Projeto

O projeto segue uma arquitetura em camadas:

```
todo-api/
├── cmd/
│   └── migrate/       # Ferramenta para migrações do banco de dados
├── db/
│   └── migrations/    # Migrações SQL
├── internal/
│   ├── config/        # Configurações da aplicação
│   ├── core/
│   │   └── task/      # Domínio e regras de negócio
│   ├── handler/       # Handlers HTTP
│   ├── router/        # Configuração de rotas
│   └── telemetry/     # Configuração OpenTelemetry
└── main.go            # Ponto de entrada da aplicação
```

## Tecnologias Utilizadas

### Bibliotecas Principais
- **gorilla/mux**: Router HTTP
- **lib/pq**: Driver PostgreSQL para Go
- **golang-migrate/migrate**: Migrações de banco de dados
- **google/uuid**: Geração de identificadores únicos
- **joho/godotenv**: Carregamento de variáveis de ambiente

### OpenTelemetry
- **go.opentelemetry.io/otel**: API principal do OpenTelemetry
- **go.opentelemetry.io/otel/trace**: API de traces
- **go.opentelemetry.io/otel/sdk**: Implementação do SDK
- **go.opentelemetry.io/otel/exporters/otlp**: Exportadores para OTLP
- **go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc**: Exportador gRPC

## Funcionalidades

A API implementa operações CRUD para tarefas:

- **GET /tasks**: Lista todas as tarefas
- **POST /task**: Cria uma nova tarefa
- **PUT /task/{id}**: Atualiza uma tarefa existente
- **DELETE /task/{id}**: Remove uma tarefa
- **GET /health**: Endpoint de verificação de saúde

## Implementação do OpenTelemetry

A telemetria foi implementada em várias camadas da aplicação:

### 1. Inicialização (internal/telemetry/tracing.go)

Configuração do provedor de traces, exportador OTLP e propagadores:

```go
func InitTracer(serviceName string, collectorURL string) (func(context.Context) error, error) {
    // Configuração do tracer, exportador e propagadores
    // ...
}
```

### 2. Middleware HTTP (internal/telemetry/middleware.go)

Instrumenta todas as requisições HTTP:

```go
func TracingMiddleware(next http.Handler) http.Handler {
    // Criação de spans para cada requisição HTTP
    // Propagação de contexto
    // Captura de status code e headers
    // ...
}
```

### 3. Camada de Serviço (internal/core/task/service.go)

Cada método de serviço cria spans próprios:

```go
func (s *TaskService) CreateTask(title, description string, concluded bool) error {
    // Criação de span específico para esta operação
    // Adição de atributos e metadados
    // ...
}
```

### 4. Repositório (internal/core/task/pg_repository.go)

Operações de banco de dados também são rastreadas:

```go
func (r *PgTaskRepository) GetAll() ([]Task, error) {
    // Span para operação de banco
    // Atributos como tipo de operação SQL
    // Métricas como número de linhas retornadas
    // ...
}
```

## Configuração do Ambiente

O projeto usa Docker Compose para configurar:

1. **PostgreSQL**: Banco de dados para armazenar as tarefas
2. **Jaeger**: Backend para visualização de traces

```yaml
version: '3.8'
services:
  postgres:
    # Configuração do PostgreSQL
  
  jaeger:
    # Configuração do Jaeger All-in-One
    # Expõe UI na porta 16686
```

## Como Executar

1. Clone o repositório
   ```
   git clone <repository-url>
   cd todo-api
   ```

2. Inicie os serviços
   ```
   docker-compose up -d
   ```

3. Execute as migrações
   ```
   go run cmd/migrate/main.go
   ```

4. Inicie a aplicação
   ```
   go run main.go
   ```

5. Visualize os traces em http://localhost:16686

## Observabilidade

A implementação de OpenTelemetry permite:

- **Rastreamento distribuído**: Acompanhe operações em todos os componentes
- **Métricas detalhadas**: Número de requisições, duração, erros
- **Diagnóstico de problemas**: Identifique gargalos de performance
- **Correlação de eventos**: Associe cada operação a um identificador único

## Benefícios do OpenTelemetry

- **Padrão aberto**: Não há lock-in de fornecedor
- **Instrumentação consistente**: Mesmo padrão em diferentes linguagens
- **Extensibilidade**: Suporte a múltiplos backends (Jaeger, Zipkin, etc.)
- **Baixo overhead**: Impacto mínimo na performance

## Contribuindo

Contribuições são bem-vindas! Este projeto serve como referência para implementação de telemetria em aplicações Go.
