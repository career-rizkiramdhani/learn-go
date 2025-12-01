package dto

type SignInRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type SignInUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SignInResult struct {
	User  SignInUser `json:"user"`
	Token string     `json:"token"`
}

type SignInData struct {
	Result SignInResult `json:"result"`
}

type SignInFullResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    SignInData `json:"data"`
}
