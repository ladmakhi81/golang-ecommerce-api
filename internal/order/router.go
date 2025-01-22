package order

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type OrderRouter struct {
	baseApi *echo.Group
	handler OrderHandler
	config  config.MainConfig
}

func NewOrderRouter(
	handler OrderHandler,
	config config.MainConfig,
) OrderRouter {
	return OrderRouter{
		handler: handler,
		config:  config,
	}
}

func (orderRouter *OrderRouter) SetBaseApi(baseApi *echo.Group) {
	orderRouter.baseApi = baseApi
}

func (orderRouter OrderRouter) RegisterRoutes() {
	orderApi := orderRouter.baseApi.Group("/orders")

	orderApi.Use(
		middlewares.AuthMiddleware(orderRouter.config.SecretKey),
	)

	orderApi.POST(
		"",
		orderRouter.handler.CreateOrder,
	)

	orderApi.PATCH(
		"/:orderId",
		orderRouter.handler.UpdateOrderStatus,
		middlewares.RoleMiddleware(userentity.AdminRole),
	)

	orderApi.GET(
		"/page",
		orderRouter.handler.FindOrdersPage,
		middlewares.RoleMiddleware(userentity.AdminRole),
	)
}
