package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/youruser/dexter-transport/internal/constant"
)

// BindReqJson200Resp เป็น generic function สำหรับ bind request และส่ง response กลับเป็น JSON 200
func BindReqJson200Resp[T any, R any](c *gin.Context, handler func(ctx context.Context, req T) (R, error)) {
	// สร้าง context สำหรับการทำงาน
	ctx := context.Background()

	// Bind JSON request body เป็น struct T
	// สำหรับ GET และ DELETE request มักไม่มี body จึงควรอ่านเฉพาะถ้าไม่ใช่ 2 วิธีนี้
	var req T
	if c.Request.Method != "GET" && c.Request.Method != "DELETE" {
		// พยายาม bind JSON ถ้ามี error ให้ส่ง 400
		if err := c.ShouldBindJSON(&req); err != nil {
			// หมายเหตุ: บางครั้งกรณี body ว่างแต่อยากให้ผ่าน ก็สามารถปรับ logic ตรงนี้ได้
			c.JSON(400, gin.H{"error": servicesconstant.INVALID_REQUEST_BODY_MESSAGE + ": " + err.Error()})
			return
		}
	}

	// เรียก handler ที่รับ context และ request struct
	resp, err := handler(ctx, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// ส่ง response กลับไปยัง client
	c.JSON(200, resp)
}
