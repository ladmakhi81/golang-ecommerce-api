package auth

import "github.com/labstack/echo/v4"

type AuthRouter struct {
	handler AuthHandler
	baseApi *echo.Group
}

func NewAuthRouter(authHandler AuthHandler) AuthRouter {
	return AuthRouter{
		handler: authHandler,
	}
}

func (authRouter *AuthRouter) SetBaseApi(baseApi *echo.Group) {
	authRouter.baseApi = baseApi
}

func (authRouter AuthRouter) RegisterRoutes() {
	authApi := authRouter.baseApi.Group("/auth")

	authApi.POST("/signup", authRouter.handler.Signup)
	authApi.POST("/login", authRouter.handler.Login)
}
