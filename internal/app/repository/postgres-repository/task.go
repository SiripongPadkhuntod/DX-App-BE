package postgres_repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/youruser/dexter-transport/internal/app/domain"
	"github.com/youruser/dexter-transport/internal/app/repository/postgres-repository/entity"
	servicesconstant "github.com/youruser/dexter-transport/internal/constant"
)

func (r *postgresRepository) CreateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	var task entity.Task
	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO ")
	queryBuilder.WriteString(servicesconstant.TABLE_TASK)
	queryBuilder.WriteString(" (title, description) VALUES (:title, :description) RETURNING id, title, description, status, created_at, updated_at")

	err := r.db.QueryRowContext(ctx, queryBuilder.String(),
		sql.Named("title", t.Title),
		sql.Named("description", t.Description),
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return mapTaskEntityToDomain(task), nil
}

func (r *postgresRepository) GetTaskByID(ctx context.Context, id int) (*domain.Task, error) {
	var task entity.Task
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, title, description, status, created_at, updated_at FROM ")
	queryBuilder.WriteString(servicesconstant.TABLE_TASK)
	queryBuilder.WriteString(" WHERE id = :id")

	err := r.db.QueryRowContext(ctx, queryBuilder.String(),
		sql.Named("id", id),
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return mapTaskEntityToDomain(task), nil
}

func (r *postgresRepository) ListTasks(ctx context.Context) ([]domain.Task, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, title, description, status, created_at, updated_at FROM ")
	queryBuilder.WriteString(servicesconstant.TABLE_TASK)
	queryBuilder.WriteString(" ORDER BY created_at DESC")

	rows, err := r.db.QueryContext(ctx, queryBuilder.String())
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
	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE ")
	queryBuilder.WriteString(servicesconstant.TABLE_TASK)
	queryBuilder.WriteString(" SET title = :title, description = :description, status = :status, updated_at = :updated_at WHERE id = :id RETURNING id, title, description, status, created_at, updated_at")

	err := r.db.QueryRowContext(ctx, queryBuilder.String(),
		sql.Named("title", t.Title),
		sql.Named("description", t.Description),
		sql.Named("status", t.Status),
		sql.Named("updated_at", time.Now()),
		sql.Named("id", t.ID),
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return mapTaskEntityToDomain(task), nil
}

func (r *postgresRepository) DeleteTask(ctx context.Context, id int) error {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("DELETE FROM ")
	queryBuilder.WriteString(servicesconstant.TABLE_TASK)
	queryBuilder.WriteString(" WHERE id = :id")

	_, err := r.db.ExecContext(ctx, queryBuilder.String(), sql.Named("id", id))
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
