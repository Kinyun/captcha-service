package models

type CommonResponse struct {
	Error        bool        `json:"error"`
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
}
