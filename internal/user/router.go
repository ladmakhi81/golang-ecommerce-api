package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type UserRouter struct {
	handler UserHandler
	baseApi *echo.Group
	config  config.MainConfig
}

func NewUserRouter(
	userHandler UserHandler,
	config config.MainConfig,
) UserRouter {
	return UserRouter{
		handler: userHandler,
		config:  config,
	}
}

func (userRouter *UserRouter) SetBaseApi(baseApi *echo.Group) {
	userRouter.baseApi = baseApi
}

func (userRouter UserRouter) RegisterRoutes() {
	usersApi := userRouter.baseApi.Group("/users")

	usersApi.Use(
		middlewares.AuthMiddleware(
			userRouter.config.SecretKey,
		),
	)

	usersApi.PATCH(
		"/verify-account/:id",
		userRouter.handler.VerifyAccountByAdmin,
		middlewares.RoleMiddleware(userentity.AdminRole),
	)

	usersApi.PATCH(
		"/complete-profile",
		userRouter.handler.CompleteProfile,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)

	usersApi.PATCH(
		"/address/active",
		userRouter.handler.AssignActiveAddressUser,
	)

	usersApi.POST(
		"/address",
		userRouter.handler.CreateUserAddress,
	)

	usersApi.GET(
		"/addresses",
		userRouter.handler.GetUserAddresses,
	)
}
