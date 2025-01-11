package auth

import (
	"github.com/labstack/echo/v4"
	authservice "github.com/ladmakhi81/golang-ecommerce-api/internal/auth/service"
)

type AuthRouter struct {
	apiRoute    *echo.Group
	handler     AuthHandler
	authService authservice.IAuthService
}

func NewAuthRouter(apiRoute *echo.Group, authService authservice.IAuthService) AuthRouter {
	return AuthRouter{
		apiRoute: apiRoute,
		handler:  NewAuthHandler(authService),
	}
}

func (router AuthRouter) SetupRouter() {
	authRouter := router.apiRoute.Group("/auth")

	authRouter.POST("/signup", router.handler.Signup)
	authRouter.POST("/login", router.handler.Login)
}
