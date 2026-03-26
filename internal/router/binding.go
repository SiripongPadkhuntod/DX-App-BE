// Package router จัดการการ binding request และการสร้าง response สำหรับ API
// ไฟล์นี้เป็น middleware layer ที่อยู่ระหว่าง HTTP handler (Gin) กับ business logic (domain/service)
//
// Flow การทำงาน:
//   HTTP Request → bindRequest (URI/Query/Body) → injectHeaders → handler(ctx, req) → JSON Response
//
// Error Handling:
//   - Bind Error        → 400 Bad Request
//   - BusinessError     → 400/401/403 (ตาม error code)
//   - TechnicalError    → 500 Internal Server Error
//   - CustomError       → HTTP status ที่กำหนดใน error
package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	servicesconstant "github.com/youruser/dexter-transport/internal/constant"
	errormodel "github.com/youruser/dexter-transport/pkg/model/errormodel"
)

// ─────────────────────────────────────────────────────────────
// Response Models — โครงสร้าง JSON ที่ส่งกลับไปยัง client
// ─────────────────────────────────────────────────────────────

// APIError คือ response body สำหรับกรณี error ทุกประเภท
// ตัวอย่าง JSON: {"code":"0001","description":"Invalid request body"}
type APIError struct {
	Code        string `json:"code"`        // รหัส error ภายใน (เช่น "0001", "0002")
	Description string `json:"description"` // คำอธิบาย error ที่ส่งให้ client
	HTTPStatus  int    `json:"-"`           // HTTP status code (ไม่ส่งใน JSON body)
}

// APIResponse คือ response body สำหรับกรณีสำเร็จ (generic type T คือ data ที่ส่งกลับ)
// ตัวอย่าง JSON: {"code":"0000","description":"success","data":{...}}
type APIResponse[T any] struct {
	Code        string `json:"code"`        // รหัสสำเร็จ ("0000")
	Description string `json:"description"` // ข้อความสำเร็จ ("success")
	Data        T      `json:"data"`        // ข้อมูล response (generic type)
	HTTPStatus  int    `json:"-"`           // HTTP status code (ไม่ส่งใน JSON body)
}

// ─────────────────────────────────────────────────────────────
// Main Entry Point — จุดเริ่มต้นหลักสำหรับทุก API endpoint
// ─────────────────────────────────────────────────────────────

// BindReqJson200Resp เป็น generic function หลักที่ใช้กับทุก API endpoint
// ทำหน้าที่ 3 อย่าง:
//  1. Bind request — อ่านข้อมูลจาก URI, Query string, และ JSON body แล้ว map เข้า struct T
//  2. Inject headers — นำ HTTP headers ใส่เข้า context เพื่อส่งต่อให้ handler ใช้งาน
//  3. Call handler & respond — เรียก business logic แล้วส่ง JSON response กลับ
//
// Type Parameters:
//   - T: ประเภทของ request struct (เช่น CreateTaskRequest)
//   - R: ประเภทของ response struct (เช่น TaskResponse)
//
// Usage (ใน router.go):
//
//	v1.POST("/tasks", func(c *gin.Context) { BindReqJson200Resp(c, h.CreateTask(c)) })
func BindReqJson200Resp[T any, R any](c *gin.Context, handler func(ctx context.Context, req T) (R, error)) {

	// ─── Step 1: Bind Request ───
	// อ่านข้อมูลจาก URI params, query string, และ JSON body แล้ว map เข้า struct
	var req T
	if err := bindRequest(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, apiErrorFromBind(err))
		return
	}

	// ─── Step 2: สร้าง context พร้อม HTTP headers ───
	// ดึง headers ทั้งหมดจาก request ใส่เข้า context.Context
	// เพื่อให้ service/repository layer สามารถเข้าถึง header ได้ (เช่น Authorization, X-Request-Id)
	ctx := injectHeaders(c)

	// ─── Step 3: เรียก business logic ───
	resp, err := handler(ctx, req)
	if err != nil {
		handleError(c, err)
		return
	}

	// ─── Step 4: ส่ง success response ───
	c.JSON(http.StatusOK, APIResponse[R]{
		Code:        servicesconstant.SUCCESS_CODE,
		Description: string(servicesconstant.SUCCESS_MESSAGE_SUCCESS),
		Data:        resp,
	})
}

// ─────────────────────────────────────────────────────────────
// Request Binding — อ่านข้อมูลจาก HTTP request เข้า struct
// ─────────────────────────────────────────────────────────────

// bindRequest ทำการ bind ข้อมูลจาก 3 แหล่ง เข้า struct เดียวกัน:
//  1. URI params   — เช่น /tasks/:id → struct field ที่มี tag `uri:"id"`
//  2. Query string — เช่น ?page=1   → struct field ที่มี tag `form:"page"`
//  3. JSON body    — เช่น {"title":"..."} → struct field ที่มี tag `json:"title"`
//
// หมายเหตุ: สำหรับ GET/DELETE จะข้าม JSON body binding เพราะปกติไม่มี body
// Validation error จาก URI/Query จะถูก ignore ชั่วคราว เพราะ field อาจถูก fill จากขั้นตอนถัดไป
func bindRequest[T any](c *gin.Context, req *T) error {
	// 1) Bind URI params (เช่น /tasks/:id)
	// ignore validation error เพราะ field อื่น ๆ อาจยังไม่ได้ bind
	if err := c.ShouldBindUri(req); err != nil {
		if _, ok := err.(validator.ValidationErrors); !ok {
			return err // error ที่ไม่ใช่ validation (เช่น type mismatch) → return ทันที
		}
	}

	// 2) Bind query string params (เช่น ?page=1&limit=10)
	if err := c.ShouldBindQuery(req); err != nil {
		if _, ok := err.(validator.ValidationErrors); !ok {
			return err
		}
	}

	// 3) Bind JSON body (เฉพาะ method ที่มี body: POST, PUT, PATCH)
	if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
		if c.Request.Body != nil && c.Request.ContentLength != 0 {
			return c.ShouldBindJSON(req) // bind JSON + validate ทั้งหมดในครั้งเดียว
		}
		// กรณี POST/PUT แต่ไม่มี body (เช่น ส่งแค่ URI params) → validate struct ด้วยตนเอง
		return validateStruct(req)
	}

	// สำหรับ GET/DELETE → validate struct ด้วยตนเอง (เพราะไม่มี JSON body binding)
	return validateStruct(req)
}

// validateStruct ใช้ validator engine ของ Gin ในการ validate struct ด้วยตนเอง
// ใช้ในกรณีที่ไม่มี JSON body binding (GET/DELETE) เพื่อให้ tag เช่น `binding:"required"` ยังทำงาน
func validateStruct(obj any) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.Struct(obj)
	}
	return nil
}

// ─────────────────────────────────────────────────────────────
// Error Handling — จัดการ error จาก business logic แล้วส่ง response
// ─────────────────────────────────────────────────────────────

// handleError แปลง error จาก domain/service layer เป็น HTTP response
// รองรับ error 3 ประเภท ตามลำดับความสำคัญ:
//
//  1. BusinessError   → error ทางธุรกิจ (เช่น ข้อมูลซ้ำ, ไม่มีสิทธิ์)
//     HTTP status ขึ้นอยู่กับ error code: 400/401/403
//
//  2. TechnicalError  → error ทางเทคนิค (เช่น DB connection, external service)
//     HTTP status: 500 เสมอ
//
//  3. CustomError     → error ที่กำหนด HTTP status เอง
//     HTTP status: ตาม error.Status ที่กำหนดไว้
func handleError(c *gin.Context, err error) {
	// BusinessError → 400/401/403 ตาม error code
	if be, ok := err.(*errormodel.BusinessError); ok {
		status := mapBusinessErrorToHTTP(be.Code)
		c.JSON(status, APIError{
			Code:        be.Code,
			Description: be.Description,
		})
		return
	}

	// TechnicalError → 500 Internal Server Error เสมอ
	// TODO: เพิ่ม logging สำหรับ stack trace เพื่อ debug
	if te, ok := err.(*errormodel.TechnicalError); ok {
		c.JSON(http.StatusInternalServerError, APIError{
			Code:        te.Code,
			Description: te.Description,
		})
		return
	}

	// CustomError → HTTP status ที่กำหนดใน error (เช่น 401, 403, 404)
	if ce, ok := err.(*errormodel.CustomError); ok {
		c.JSON(ce.Status, APIError{
			Code:        ce.Code,
			Description: ce.Description,
		})
		return
	}
}

// apiErrorFromBind สร้าง APIError จาก bind/validation error
// ใช้เมื่อ request body/query/URI ไม่ถูกต้องตาม struct tag
func apiErrorFromBind(err error) APIError {
	return APIError{
		Code:        errormodel.InvalidRequest,
		Description: err.Error(),
	}
}

// ─────────────────────────────────────────────────────────────
// Context Helpers — จัดการ context สำหรับส่งต่อข้อมูลให้ handler
// ─────────────────────────────────────────────────────────────

// injectHeaders ดึง HTTP headers ทั้งหมดจาก request แล้วใส่เข้า context.Context
// ทำให้ service/repository layer สามารถเข้าถึง header ได้ผ่าน ctx.Value(key)
//
// ตัวอย่างการใช้ใน service layer:
//
//	authToken := ctx.Value("Authorization").(string)
func injectHeaders(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	for key, values := range c.Request.Header {
		if len(values) > 0 {
			ctx = context.WithValue(ctx, key, values[0])
		}
	}
	return ctx
}

// ─────────────────────────────────────────────────────────────
// HTTP Status Mapping — กำหนด HTTP status ตาม error code
// ─────────────────────────────────────────────────────────────

// mapBusinessErrorToHTTP แปลง business error code เป็น HTTP status code
//
//	"0002" (Unauthorized) → 401
//	"0003" (Forbidden)    → 403
//	อื่น ๆ                → 400 (Bad Request)
func mapBusinessErrorToHTTP(code string) int {
	switch code {
	case errormodel.Unauthorized:
		return http.StatusUnauthorized
	case errormodel.Forbidden:
		return http.StatusForbidden
	default:
		return http.StatusBadRequest
	}
}
