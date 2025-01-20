package order

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
)

type OrderRouter struct {
	apiRouter    *echo.Group
	orderHandler OrderHandler
	config       config.MainConfig
	orderService orderservice.IOrderService
}

func NewOrderRouter(
	apiRouter *echo.Group,
	config config.MainConfig,
	orderService orderservice.IOrderService,
) OrderRouter {
	return OrderRouter{
		apiRouter:    apiRouter,
		orderHandler: NewOrderHandler(orderService),
		config:       config,
	}
}

func (orderRouter OrderRouter) SetupRouter() {
	orderApi := orderRouter.apiRouter.Group("/orders")

	orderApi.Use(
		middlewares.AuthMiddleware(orderRouter.config.SecretKey),
	)

	orderApi.POST(
		"",
		orderRouter.orderHandler.CreateOrder,
	)

	orderApi.PATCH(
		"/:orderId",
		orderRouter.orderHandler.UpdateOrderStatus,
	)

	orderApi.GET(
		"/page",
		orderRouter.orderHandler.FindOrdersPage,
	)
}
