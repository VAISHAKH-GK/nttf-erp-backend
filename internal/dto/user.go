package dto

// LoginReq represents login request body
//
//	@Description	Login request payload
type LoginReq struct {
	Email    string `json:"email" example:"admin" validate:"required" minLength:"3" maxLength:"50"`
	Password string `json:"password" example:"password" validate:"required,min=6" minLength:"6"`
}

// LoginRes represents login response
//
//	@Description	Login response with tokens
type LoginRes struct {
	AuthToken    string `json:"authToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refreshToken" example:"550e8400-e29b-41d4-a716-446655440000"`
}
