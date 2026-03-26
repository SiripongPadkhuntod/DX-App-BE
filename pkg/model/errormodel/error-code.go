package errormodel

const (
	InvalidRequest    = "0001"
	Unauthorized      = "0002"
	Forbidden         = "0003"
	InvalidRequestCdp = "5101"
	InvalidJwt        = "0004"
	AlreadyExists     = "0005"
	GenericErrorCode  = "9999"
)

const (
	ForbiddenMsgTh        = "รหัสผู้ใช้, รหัสพิน, หรือโทเคนรีเฟรชไม่ถูกต้อง"
	InvalidJwtMsgTh       = "ไม่สามารถตรวจสอบหรือแยกวิเคราะห์ JWT ได้"
	GenericErrorCodeMsgTh = "ข้อผิดพลาดเซิร์ฟเวอร์ภายใน"
	AlreadyExistsMsgTh    = "ข้อมูลนี้มีอยู่แล้วในระบบ"
	InvalidRequestTh      = "ข้อมูลไม่ถูกต้อง"
)
const (
	ForbiddenMsgEn        = "Invalid user ID, pin or refresh token"
	InvalidJwtMsgEn       = "Unable to validate or parse jwt"
	GenericErrorCodeMsgEn = "Internal Server Error"
	AlreadyExistsMsgEn    = "This data already exists in the system"
	InvalidRequestEn      = "Invalid Request"
)
