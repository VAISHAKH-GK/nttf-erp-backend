package dto

type ErrorRes struct {
	Error   string `json:"error" example:"Invalid credentials"`
	Message string `json:"message,omitempty" example:"Email or password is incorrect"`
}
