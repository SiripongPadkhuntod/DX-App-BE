package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	handlerdto "github.com/youruser/dexter-transport/internal/app/handler/dto"
	servicesconstant "github.com/youruser/dexter-transport/internal/constant"
	"github.com/youruser/dexter-transport/pkg/v1/dto"
)

// HealthCheck godoc
// @Summary      Show service health status
// @Description  Get the health status of the service
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  handlerdto.HealthResponse
// @Router       /api/v1/health [get]
func (h *handler) Health(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.HealthResponse, error) {
	return func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.HealthResponse, error) {
		// เรียก service layer เพื่อทำการตรวจสอบสุขภาพของระบบ
		if h.svc != nil {
			id, serviceName, err := h.svc.Health(ctx)
			if err != nil {
				// ถ้ามี error เกิดขึ้น ให้ส่ง response กลับไปว่าไม่ผ่าน
				return &handlerdto.HealthResponse{
					Status:  servicesconstant.UNHEALTHY_STATUS,
					Code:    servicesconstant.DATABASE_ERROR_CODE,
					Message: err.Error(),
				}, nil
			}

			// ถ้าทุกอย่างปกติ ให้ส่ง response กลับไปว่า healthy พร้อมข้อมูลจาก DB
			return &handlerdto.HealthResponse{
				Status:      string(servicesconstant.HEALTHY_STATUS),
				Code:        servicesconstant.SUCCESS_CODE,
				Message:     string(servicesconstant.HEALTHY_MESSAGE),
				HealthId:    id,
				ServiceName: serviceName,
			}, nil
		}

		return &handlerdto.HealthResponse{
			Status:  string(servicesconstant.HEALTHY_STATUS),
			Code:    servicesconstant.SUCCESS_CODE,
			Message: string(servicesconstant.HEALTHY_MESSAGE),
		}, nil
	}
}
