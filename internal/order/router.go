package order

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type OrderRouter struct {
	baseApi    *echo.Group
	handler    OrderHandler
	middleware middlewares.Middleware
}

func NewOrderRouter(
	handler OrderHandler,
	middleware middlewares.Middleware,
) OrderRouter {
	return OrderRouter{
		handler:    handler,
		middleware: middleware,
	}
}

func (orderRouter *OrderRouter) SetBaseApi(baseApi *echo.Group) {
	orderRouter.baseApi = baseApi
}

func (orderRouter OrderRouter) RegisterRoutes() {
	orderApi := orderRouter.baseApi.Group("/orders")

	orderApi.Use(
		orderRouter.middleware.AuthMiddleware(),
	)

	orderApi.POST(
		"",
		orderRouter.handler.CreateOrder,
	)

	orderApi.PATCH(
		"/:orderId",
		orderRouter.handler.UpdateOrderStatus,
		orderRouter.middleware.RoleMiddleware(userentity.AdminRole),
	)

	orderApi.GET(
		"/page",
		orderRouter.handler.FindOrdersPage,
		orderRouter.middleware.RoleMiddleware(userentity.AdminRole),
	)
}
