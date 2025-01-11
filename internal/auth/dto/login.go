package authdto

type LoginReqBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewLoginResponse(accessToken string) *LoginResponse {
	return &LoginResponse{AccessToken: accessToken}
}
