package task

import (
	"context"
	"database/sql"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type PgTaskRepository struct {
	db *sql.DB
}

func NewPgTaskRepository(db *sql.DB) *PgTaskRepository {
	return &PgTaskRepository{db: db}
}

func (r *PgTaskRepository) Create(task Task) error {
	ctx := context.Background()
	tracer := otel.Tracer("postgres-repository")

	ctx, span := tracer.Start(ctx, "CreateTask", trace.WithAttributes(
		attribute.String("db.operation", "INSERT"),
		attribute.String("task.id", task.Id),
	))
	defer span.End()

	query := `INSERT INTO tasks (id, title, description, concluded) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, task.Id, task.Title, task.Description, task.Concluded)
	return err
}

func (r *PgTaskRepository) GetAll() ([]Task, error) {
	ctx := context.Background()
	tracer := otel.Tracer("postgres-repository")

	ctx, span := tracer.Start(ctx, "GetAllTasks", trace.WithAttributes(
		attribute.String("db.operation", "SELECT"),
	))
	defer span.End()

	query := `SELECT id, title, description, concluded FROM tasks`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Concluded); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	span.SetAttributes(attribute.Int("db.rows_returned", len(tasks)))
	return tasks, nil
}

func (r *PgTaskRepository) Update(id string, task Task) error {
	ctx := context.Background()
	tracer := otel.Tracer("postgres-repository")

	ctx, span := tracer.Start(ctx, "UpdateTask", trace.WithAttributes(
		attribute.String("db.operation", "UPDATE"),
		attribute.String("task.id", id),
	))
	defer span.End()

	query := `UPDATE tasks SET title = $1, description = $2, concluded = $3 WHERE id = $4`
	result, err := r.db.ExecContext(ctx, query, task.Title, task.Description, task.Concluded, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	span.SetAttributes(attribute.Int("db.rows_affected", int(rowsAffected)))

	if rowsAffected == 0 {
		return errors.New("task não encontrada")
	}
	return nil
}

func (r *PgTaskRepository) Delete(id string) error {
	ctx := context.Background()
	tracer := otel.Tracer("postgres-repository")

	ctx, span := tracer.Start(ctx, "DeleteTask", trace.WithAttributes(
		attribute.String("db.operation", "DELETE"),
		attribute.String("task.id", id),
	))
	defer span.End()

	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	span.SetAttributes(attribute.Int("db.rows_affected", int(rowsAffected)))

	if rowsAffected == 0 {
		return errors.New("task não encontrada")
	}
	return nil
}
