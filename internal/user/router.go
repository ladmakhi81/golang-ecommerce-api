package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
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
	userAddressService userservice.IUserAddressService,
	config config.MainConfig,
) UserRouter {
	return UserRouter{
		config:      config,
		apiRouter:   apiRouter,
		userHandler: NewUserHandler(userService, userAddressService),
	}
}

func (userRouter UserRouter) SetupRouter() {
	usersApi := userRouter.apiRouter.Group("/users")

	usersApi.Use(
		middlewares.AuthMiddleware(
			userRouter.config.SecretKey,
		),
	)

	usersApi.PATCH(
		"/verify-account/:id",
		userRouter.userHandler.VerifyAccountByAdmin,
		middlewares.RoleMiddleware(userentity.AdminRole),
	)

	usersApi.PATCH(
		"/complete-profile",
		userRouter.userHandler.CompleteProfile,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)

	usersApi.PATCH(
		"/address/active",
		userRouter.userHandler.AssignActiveAddressUser,
	)

	usersApi.POST(
		"/address",
		userRouter.userHandler.CreateUserAddress,
	)

	usersApi.GET(
		"/addresses",
		userRouter.userHandler.GetUserAddresses,
	)
}
