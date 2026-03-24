package handlerdto

type HealthResponse struct {
	Status      string `json:"status"`
	Code        string `json:"code"`
	Message     string `json:"message"`
	HealthId    string `json:"health_id,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
}
