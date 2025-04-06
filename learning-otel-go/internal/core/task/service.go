package task

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title, description string, concluded bool) error {
	// Obter o tracer para este service
	ctx := context.Background()
	tracer := otel.Tracer("task-service")

	// Criar um span para a operação
	ctx, span := tracer.Start(ctx, "CreateTask", trace.WithAttributes(
		attribute.String("task.title", title),
	))
	defer span.End()

	id := uuid.New().String()
	task := Task{
		Id:          id,
		Title:       title,
		Description: description,
		Concluded:   concluded,
	}

	// Adicionar ID ao span para correlacionar
	span.SetAttributes(attribute.String("task.id", id))

	return s.repo.Create(task)
}

func (s *TaskService) GetTasks() ([]Task, error) {
	ctx := context.Background()
	tracer := otel.Tracer("task-service")

	ctx, span := tracer.Start(ctx, "GetTasks")
	defer span.End()

	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	span.SetAttributes(attribute.Int("task.count", len(tasks)))
	return tasks, nil
}

func (s *TaskService) EditTask(id string, newTask Task) error {
	ctx := context.Background()
	tracer := otel.Tracer("task-service")

	ctx, span := tracer.Start(ctx, "EditTask", trace.WithAttributes(
		attribute.String("task.id", id),
	))
	defer span.End()

	return s.repo.Update(id, newTask)
}

func (s *TaskService) DeleteTask(id string) error {
	ctx := context.Background()
	tracer := otel.Tracer("task-service")

	ctx, span := tracer.Start(ctx, "DeleteTask", trace.WithAttributes(
		attribute.String("task.id", id),
	))
	defer span.End()

	return s.repo.Delete(id)
}
