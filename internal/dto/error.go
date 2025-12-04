package dto

type ErrorRes struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
