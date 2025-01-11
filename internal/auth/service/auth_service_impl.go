package authservice

import (
	"net/http"

	authdto "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/dto"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
)

type AuthService struct {
	userService  userservice.IUserService
	jwtService   IJwtService
	emailService pkgemail.IEmailService
}

func NewAuthService(
	userService userservice.IUserService,
	jwtService IJwtService,
	emailService pkgemail.IEmailService,
) AuthService {
	return AuthService{
		userService,
		jwtService,
		emailService,
	}
}

func (authService AuthService) Signup(reqBody authdto.SignupReqBody) (*authdto.SignupResponse, error) {
	if reqBody.Role == userentity.AdminRole {
		return nil, types.NewClientError("Invalid role type", http.StatusBadRequest)
	}
	createdUser, createdUserErr := authService.userService.CreateBasicUser(reqBody.Email, reqBody.Password, reqBody.Role)
	if createdUserErr != nil {
		return nil, createdUserErr
	}
	generatedAccessToken, generatedAccessTokenErr := authService.jwtService.GenerateAccessToken(createdUser)
	if generatedAccessTokenErr != nil {
		return nil, types.NewServerError(
			"error in generate access token",
			"authservice.signup.jwtService.generateAccessToken",
			generatedAccessTokenErr,
		)
	}
	authService.emailService.SendEmail(
		pkgemaildto.NewSendEmailDto(
			createdUser.Email,
			"Your Account Created Successfully",
			"Welcome To GoEcommerce Application",
		),
	)
	return authdto.NewSignupResponse(generatedAccessToken), nil
}

func (authService AuthService) Login(reqBody authdto.LoginReqBody) (*authdto.LoginResponse, error) {
	user, err := authService.userService.FindUserByEmailAndPassword(reqBody.Email, reqBody.Password)
	if err != nil {
		return nil, err
	}
	accessToken, accessTokenErr := authService.jwtService.GenerateAccessToken(user)
	if accessTokenErr != nil {
		return nil, accessTokenErr
	}
	return authdto.NewLoginResponse(accessToken), nil
}
