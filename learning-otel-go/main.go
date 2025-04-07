package main

import (
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"

	"todo-api/internal/app"
)

func main() {
	// Criar contexto raiz da aplicação
	ctx := context.Background()

	// Criar e inicializar a aplicação
	application := app.New()

	if err := application.Initialize(ctx); err != nil {
		log.Fatalf("Falha ao inicializar aplicação: %v", err)
		os.Exit(1)
	}

	// Iniciar o servidor
	if err := application.Start(ctx); err != nil {
		log.Fatalf("Erro fatal na aplicação: %v", err)
		os.Exit(1)
	}
}
