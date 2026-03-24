package postgres_repository

import (
	"context"
	"time"

	"github.com/youruser/dexter-transport/internal/app/domain"
	"github.com/youruser/dexter-transport/internal/app/repository/postgres-repository/entity"
)

func (r *postgresRepository) CreateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	var task entity.Task
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id, title, description, status, created_at, updated_at",
		t.Title, t.Description,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return mapTaskEntityToDomain(task), nil
}

func (r *postgresRepository) GetTaskByID(ctx context.Context, id int) (*domain.Task, error) {
	var task entity.Task
	err := r.db.QueryRowContext(ctx,
		"SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1",
		id,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return mapTaskEntityToDomain(task), nil
}

func (r *postgresRepository) ListTasks(ctx context.Context) ([]domain.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, *mapTaskEntityToDomain(task))
	}
	return tasks, nil
}

func (r *postgresRepository) UpdateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	var task entity.Task
	err := r.db.QueryRowContext(ctx,
		"UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5 RETURNING id, title, description, status, created_at, updated_at",
		t.Title, t.Description, t.Status, time.Now(), t.ID,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return mapTaskEntityToDomain(task), nil
}

func (r *postgresRepository) DeleteTask(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tasks WHERE id = $1", id)
	return err
}

func mapTaskEntityToDomain(e entity.Task) *domain.Task {
	return &domain.Task{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Status:      e.Status,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
