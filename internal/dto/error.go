package dto

type ErrorRes struct {
	Error   string `json:"error" example:"Invalid credentials"`
	Message string `json:"message,omitempty" example:"Username or password is incorrect"`
}
