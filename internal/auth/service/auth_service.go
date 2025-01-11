package authservice

import authdto "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/dto"

type IAuthService interface {
	Signup(reqBody authdto.SignupReqBody) (*authdto.SignupResponse, error)
	Login(reqBody authdto.LoginReqBody) (*authdto.LoginResponse, error)
}
