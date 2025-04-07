package handler

import (
	"encoding/json"
	"net/http"

	"todo-api/internal/core/task"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	service *task.TaskService
}

func NewTaskHandler(service *task.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskBody task.Task
	err := json.NewDecoder(r.Body).Decode(&taskBody)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	err = h.service.CreateTask(taskBody.Title, taskBody.Description, taskBody.Concluded)
	if err != nil {
		http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		http.Error(w, "Erro ao buscar tarefas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var taskBody task.Task
	err := json.NewDecoder(r.Body).Decode(&taskBody)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	err = h.service.EditTask(id, taskBody)
	if err != nil {
		http.Error(w, "Erro ao editar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, "Erro ao deletar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
