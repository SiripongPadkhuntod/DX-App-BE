package modelV1

import "time"

type Health struct {
	HealthId   string    `json:"health_id"`
	Service    string    `json:"service"`
	Status     int       `json:"status"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
