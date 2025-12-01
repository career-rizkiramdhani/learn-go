package dto

type SignInRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
