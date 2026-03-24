package handlerdto

// สำหรับการลบมักไม่ใช้ body แต่สามารถใส่ struct ว่างไว้เพื่อความสอดคล้อง
type DeleteTaskRequest struct{}

type DeleteTaskResponse struct {
	Message string `json:"message"`
}