package entity

import "time"

type Task struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
