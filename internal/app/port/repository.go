package port

import (
	"context"

	"github.com/youruser/dexter-transport/internal/app/domain"
)

type Repository struct {
	Sql           SqlRepository
	Misc          MiscRepository
	ObjectStorage ObjectStorageRepository
	Cache         CacheRepository
}

type SqlRepository interface {
	Ping(ctx context.Context) error
	GetFirstHealthRecord(ctx context.Context) (*domain.Health, error)

	// Task Domain methods
	CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error)
	GetTaskByID(ctx context.Context, id int) (*domain.Task, error)
	ListTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) (*domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
}
type MiscRepository interface {
	// GetCurrentDate() time.Time

	// NewUUID() string
}

type ObjectStorageRepository interface {
	// UploadObject(ctx context.Context, objectName string, body io.ReadCloser, contentLength int64, contentType string) (string, error)
	// DeleteObject(ctx context.Context, objectName string) error
	// ObjectURL(ctx context.Context, objectName string) (string, error)
	// GenerateUploadPresignedURL(ctx context.Context, objectName string) (string, error)
	// BasePath() string
}

type CacheRepository interface {
	// Ping(ctx context.Context) error
	// Set(ctx context.Context, key, value string, ttl time.Duration) error
	// SetNX(ctx context.Context, key, value string, ttl time.Duration) (bool, error)
	// Get(ctx context.Context, key string) (string, error)
	// Delete(ctx context.Context, key string) (bool, error)
}
