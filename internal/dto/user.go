package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}
