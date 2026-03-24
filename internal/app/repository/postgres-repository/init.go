package postgres_repository

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/youruser/dexter-transport/internal/app/port"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) port.SqlRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}
