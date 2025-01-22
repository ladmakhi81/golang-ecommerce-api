package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type UserRouter struct {
	handler    UserHandler
	baseApi    *echo.Group
	middleware middlewares.Middleware
}

func NewUserRouter(
	userHandler UserHandler,
	middleware middlewares.Middleware,
) UserRouter {
	return UserRouter{
		handler:    userHandler,
		middleware: middleware,
	}
}

func (userRouter *UserRouter) SetBaseApi(baseApi *echo.Group) {
	userRouter.baseApi = baseApi
}

func (userRouter UserRouter) RegisterRoutes() {
	usersApi := userRouter.baseApi.Group("/users")

	usersApi.Use(
		userRouter.middleware.AuthMiddleware(),
	)

	usersApi.PATCH(
		"/verify-account/:id",
		userRouter.handler.VerifyAccountByAdmin,
		userRouter.middleware.RoleMiddleware(userentity.AdminRole),
	)

	usersApi.PATCH(
		"/complete-profile",
		userRouter.handler.CompleteProfile,
		userRouter.middleware.RoleMiddleware(userentity.VendorRole),
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
