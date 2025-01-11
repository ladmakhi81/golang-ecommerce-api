package authdto

import user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"

type SignupReqBody struct {
	Role     user_entity.UserRole `json:"role" validate:"required"`
	Email    string               `json:"email" validate:"required,email"`
	Password string               `json:"password" validate:"required,min=8"`
}

type SignupResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewSignupResponse(accessToken string) *SignupResponse {
	return &SignupResponse{AccessToken: accessToken}
}
