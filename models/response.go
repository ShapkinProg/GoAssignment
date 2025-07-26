package models

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
	}
}

func Error(message string) APIResponse {
	return APIResponse{
		Success: false,
		Error:   message,
	}
}
