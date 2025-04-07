package router

import (
	"net/http"

	"todo-api/internal/core/task"
	"todo-api/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(service *task.TaskService) http.Handler {
	r := mux.NewRouter()

	taskHandler := handler.NewTaskHandler(service)

	// Configurar rotas
	r.HandleFunc("/task", taskHandler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", taskHandler.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", taskHandler.EditTaskHandler).Methods("PUT")
	r.HandleFunc("/task/{id}", taskHandler.DeleteTaskHandler).Methods("DELETE")

	// Adicionar health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return r
}
