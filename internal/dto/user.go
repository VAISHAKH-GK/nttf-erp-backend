package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}
