package service

import (
	"github.com/youruser/dexter-transport/internal/app/port"
)

// service คือ struct ที่จะ implement port.Service interface
// ใช้ตัวพิมพ์เล็ก (unexported) เพื่อบังคับให้สร้างผ่าน New() เท่านั้น
type service struct {
	repo port.Repository
}

// New เป็น Constructor สำหรับสร้าง Service
// โดยรับ port.Repository เข้ามา (Dependency Injection)
func New(repo port.Repository) port.Service {
	return &service{repo: repo}
}