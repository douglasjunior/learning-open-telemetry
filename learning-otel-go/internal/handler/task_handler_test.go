package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-api/internal/core/task"

	"github.com/gorilla/mux"
)

// Função para simular variáveis de rota do Mux
func muxSetVars(r *http.Request, vars map[string]string) *http.Request {
	return mux.SetURLVars(r, vars)
}

func setupMockHandler() *TaskHandler {
	mockRepo := task.NewMockTaskRepository()
	service := task.NewTaskService(mockRepo)

	service.CreateTask("Teste", "Descrição", false)

	return NewTaskHandler(service)
}

func TestCreateTaskHandler(t *testing.T) {
	handler := setupMockHandler()

	body := map[string]interface{}{
		"title":       "Nova Task",
		"description": "Descrição da nova task",
		"concluded":   false,
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/task", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateTaskHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Esperava status 201, obtive %d", status)
	}
}

func TestGetTasksHandler(t *testing.T) {
	handler := setupMockHandler()

	req := httptest.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()

	handler.GetTasksHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Esperava 200, obtive %d", rr.Code)
	}

	var tasks []task.Task
	err := json.NewDecoder(rr.Body).Decode(&tasks)
	if err != nil {
		t.Errorf("Erro ao decodificar JSON: %v", err)
	}

	if len(tasks) == 0 {
		t.Errorf("Esperava ao menos 1 tarefa")
	}
}

func TestEditTaskHandler(t *testing.T) {
	mockRepo := task.NewMockTaskRepository()
	svc := task.NewTaskService(mockRepo)
	handler := NewTaskHandler(svc)

	taskData := task.Task{
		Id:          "123",
		Title:       "Antigo",
		Description: "Desc",
		Concluded:   false,
	}
	mockRepo.Create(taskData)

	updated := task.Task{
		Title:       "Atualizado",
		Description: "Nova Desc",
		Concluded:   true,
	}
	body, _ := json.Marshal(updated)

	req := httptest.NewRequest("PUT", "/task/123", bytes.NewBuffer(body))
	req = muxSetVars(req, map[string]string{"id": "123"})
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.EditTaskHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Esperava 200, obtive %d", rr.Code)
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	mockRepo := task.NewMockTaskRepository()
	svc := task.NewTaskService(mockRepo)
	handler := NewTaskHandler(svc)

	mockRepo.Create(task.Task{
		Id:          "abc",
		Title:       "Apagar",
		Description: "Desc",
		Concluded:   false,
	})

	req := httptest.NewRequest("DELETE", "/task/abc", nil)
	req = muxSetVars(req, map[string]string{"id": "abc"})
	rr := httptest.NewRecorder()

	handler.DeleteTaskHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Esperava 200, obtive %d", rr.Code)
	}
}
