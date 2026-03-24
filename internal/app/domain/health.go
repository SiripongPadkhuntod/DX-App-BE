package domain

import "time"

type Health struct {
	HealthId   string
	Service    string
	Status     int
	CreateDate time.Time
	UpdateDate time.Time
}
