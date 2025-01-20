package authservice

import (
	"net/http"

	authdto "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/dto"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
	pkgemaildto "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/dto"
	pkgemail "github.com/ladmakhi81/golang-ecommerce-api/pkg/email/service"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type AuthService struct {
	userService  userservice.IUserService
	jwtService   IJwtService
	emailService pkgemail.IEmailService
	translation  translations.ITranslation
}

func NewAuthService(
	userService userservice.IUserService,
	jwtService IJwtService,
	emailService pkgemail.IEmailService,
	translation translations.ITranslation,
) AuthService {
	return AuthService{
		userService:  userService,
		jwtService:   jwtService,
		emailService: emailService,
		translation:  translation,
	}
}

func (authService AuthService) Signup(reqBody authdto.SignupReqBody) (*authdto.SignupResponse, error) {
	if reqBody.Role == userentity.AdminRole {
		return nil, types.NewClientError(
			authService.translation.Message("auth.invalid_role"),
			http.StatusBadRequest,
		)
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
			authService.translation.Message("auth.signup_subject_email"),
			authService.translation.Message("auth.signup_body_email"),
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
