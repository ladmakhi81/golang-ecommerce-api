package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authdto "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/dto"
	authservice "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
)

type AuthHandler struct {
	AuthService authservice.IAuthService
}

func NewAuthHandler(authService authservice.IAuthService) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

func (handler AuthHandler) Signup(c echo.Context) error {
	var reqBody authdto.SignupReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	res, err := handler.AuthService.Signup(reqBody)
	if err != nil {
		return err
	}
	c.JSON(
		http.StatusCreated,
		map[string]any{
			"success": true,
			"data":    res,
		},
	)
	return nil
}

func (handler AuthHandler) Login(c echo.Context) error {
	var reqBody authdto.LoginReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	res, err := handler.AuthService.Login(reqBody)
	if err != nil {
		return err
	}
	c.JSON(
		http.StatusOK,
		map[string]any{
			"success": true,
			"data":    res,
		},
	)
	return nil
}
