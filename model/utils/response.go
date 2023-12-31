package utils

// Response represents the custom response format.
type Response struct {
	Code    int               `json:"code"`
	Message map[string]string `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
}

// NewResponse creates a new Response instance.
func NewResponse(code int, message map[string]string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse creates an error Response instance.
func ErrorResponse(code int, message map[string]string) *Response {
	return NewResponse(code, message, nil)
}

// SuccessResponse creates a success Response instance.
func SuccessResponse(code int, message map[string]string, data interface{}) *Response {
	return NewResponse(code, message, data)
}
