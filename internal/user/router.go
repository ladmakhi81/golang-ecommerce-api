package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userservice "github.com/ladmakhi81/golang-ecommerce-api/internal/user/service"
)

type UserRouter struct {
	apiRouter   *echo.Group
	userHandler UserHandler
	config      config.MainConfig
}

func NewUserRouter(
	apiRouter *echo.Group,
	userService userservice.IUserService,
	config config.MainConfig,
) UserRouter {
	return UserRouter{
		config:      config,
		apiRouter:   apiRouter,
		userHandler: NewUserHandler(userService),
	}
}

func (userRouter UserRouter) SetupRouter() {
	usersRouter := userRouter.apiRouter.Group("/users")

	usersRouter.PATCH(
		"/verify-account/{id}",
		userRouter.userHandler.VerifyAccountByAdmin,
		middlewares.AuthMiddleware(
			userRouter.config.SecretKey,
		),
	)

	usersRouter.PATCH(
		"/complete-profile",
		userRouter.userHandler.CompleteProfile,
		middlewares.AuthMiddleware(
			userRouter.config.SecretKey,
		),
	)
}
