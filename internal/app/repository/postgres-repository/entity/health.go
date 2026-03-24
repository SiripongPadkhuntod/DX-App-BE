package entity

import "time"

type Health struct {
	HealthId   string    `db:"health_id"`
	Service    string    `db:"service"`
	Status     int       `db:"status"`
	CreateDate time.Time `db:"create_date"`
	UpdateDate time.Time `db:"update_date"`
}
