package service

import (
	"context"
	"log"

)

func (s *service) Health(ctx context.Context) (id string, service string, err error) {
	// 1. ตรวจสอบสถานะการเชื่อมต่อกับ Database
	if err = s.repo.Sql.Ping(ctx); err != nil {
		log.Printf("Health check failed (Ping): %v", err)
		return "", "", err
	}

	// 2. ดึงข้อมูลจาก database ผ่าน domain model
	h, err := s.repo.Sql.GetFirstHealthRecord(ctx)
	if err != nil {
		log.Printf("Health check failed (Query): %v", err)
		return "", "", err
	}

	log.Println("Health check successful")
	return h.HealthId, h.Service, nil
}