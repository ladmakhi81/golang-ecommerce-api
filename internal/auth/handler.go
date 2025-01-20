package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	authdto "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/dto"
	authservice "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/service"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/pkg/translations"
)

type AuthHandler struct {
	AuthService authservice.IAuthService
	translation translations.ITranslation
}

func NewAuthHandler(
	authService authservice.IAuthService,
	translation translations.ITranslation,
) AuthHandler {
	return AuthHandler{
		AuthService: authService,
		translation: translation,
	}
}

func (handler AuthHandler) Signup(c echo.Context) error {
	var reqBody authdto.SignupReqBody
	if err := c.Bind(&reqBody); err != nil {
		fmt.Println("error")
		return types.NewClientError(
			handler.translation.Message("errors.invalid_request_body"),
			http.StatusBadRequest,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	token, err := handler.AuthService.Signup(reqBody)
	if err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusCreated,
		token,
	)
	return nil
}

func (handler AuthHandler) Login(c echo.Context) error {
	var reqBody authdto.LoginReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError(
			handler.translation.Message("errors.invalid_request_body"),
			http.StatusBadRequest,
		)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	token, err := handler.AuthService.Login(reqBody)
	if err != nil {
		return err
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		token,
	)
	return nil
}
