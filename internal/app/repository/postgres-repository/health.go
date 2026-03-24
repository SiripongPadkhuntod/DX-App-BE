package postgres_repository

import (
	"context"

	"github.com/youruser/dexter-transport/internal/app/domain"
	"github.com/youruser/dexter-transport/internal/app/repository/postgres-repository/entity"
)

func (r *postgresRepository) GetFirstHealthRecord(ctx context.Context) (*domain.Health, error) {
	var h entity.Health
	err := r.db.QueryRowContext(ctx, "SELECT health_id, service FROM health LIMIT 1").Scan(&h.HealthId, &h.Service)
	if err != nil {
		return nil, err
	}
	return &domain.Health{
		HealthId: h.HealthId,
		Service:  h.Service,
	}, nil
}
