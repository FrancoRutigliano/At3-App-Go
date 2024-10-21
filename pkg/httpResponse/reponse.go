package httpresponse

type ApiResponse struct {
	StatusCode int         `json:"status"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewApiError(statusCode int, msg string, data interface{}) *ApiResponse {
	return &ApiResponse{
		StatusCode: statusCode,
		Msg:        msg,
		Data:       data,
	}
}
