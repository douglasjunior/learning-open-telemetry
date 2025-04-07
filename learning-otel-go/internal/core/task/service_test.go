package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	mockRepo := NewMockTaskRepository()
	service := NewTaskService(mockRepo)

	err := service.CreateTask("Título", "Descrição", false)
	if err != nil {
		t.Errorf("Erro inesperado ao criar tarefa: %v", err)
	}

	tasks, _ := mockRepo.GetAll()
	if len(tasks) != 1 {
		t.Errorf("Esperava 1 tarefa, mas tenho %d", len(tasks))
	}
}

func TestEditTask(t *testing.T) {

	mockRepo := NewMockTaskRepository()
	service := NewTaskService(mockRepo)

	service.CreateTask("Antigo", "Desc", false)

	tasks, _ := mockRepo.GetAll()
	if len(tasks) != 1 {
		t.Fatalf("Falha ao criar tarefa inicial")
	}
	taskId := tasks[0].Id

	update := Task{Title: "Novo", Description: "Nova Desc", Concluded: true}
	service.EditTask(taskId, update)

	tasksAtualizadas, _ := mockRepo.GetAll()
	if tasksAtualizadas[0].Title != "Novo" {
		t.Errorf("Esperava título 'Novo', obtive '%s'", tasksAtualizadas[0].Title)
	}
}

func TestDeleteTask(t *testing.T) {
	mockRepo := NewMockTaskRepository()
	service := NewTaskService(mockRepo)

	task := Task{Id: "123", Title: "Excluir", Description: "Desc", Concluded: false}
	mockRepo.tasks = append(mockRepo.tasks, task)

	service.DeleteTask("123")

	tasks, _ := mockRepo.GetAll()
	if len(tasks) != 0 {
		t.Errorf("Esperava lista vazia, mas tenho %d item(ns)", len(tasks))
	}
}
