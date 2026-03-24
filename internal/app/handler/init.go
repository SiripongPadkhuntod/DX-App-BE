package handler

import (
	"github.com/youruser/dexter-transport/internal/app/port"
)

// handler คือ struct ที่จะ implement port.Handler interface
// ใช้ตัวพิมพ์เล็ก (unexported) เพื่อบังคับให้สร้างผ่าน New() เท่านั้น
type handler struct {
	svc port.Service
}

// New เป็น Constructor สำหรับสร้าง Handler
// โดยรับ port.Handler เข้ามา (Dependency Injection)

func New(svc port.Service) port.Handler {

	return &handler{svc: svc}
}
