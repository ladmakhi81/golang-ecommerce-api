package payment

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type PaymentRouter struct {
	baseApi *echo.Group
	handler PaymentHandler
	config  config.MainConfig
}

func NewPaymentRouter(
	config config.MainConfig,
	handler PaymentHandler,
) PaymentRouter {
	return PaymentRouter{
		config:  config,
		handler: handler,
	}
}

func (paymentRouter *PaymentRouter) SetBaseApi(baseApi *echo.Group) {
	paymentRouter.baseApi = baseApi
}

func (paymentRouter PaymentRouter) RegisterRoutes() {
	paymentsApi := paymentRouter.baseApi.Group("/payments")

	paymentsApi.Use(
		middlewares.AuthMiddleware(paymentRouter.config.SecretKey),
	)

	paymentsApi.POST(
		"/verify",
		paymentRouter.handler.VerifyPayment,
	)
	paymentsApi.GET(
		"/page",
		paymentRouter.handler.GetPaymentsPage,
		middlewares.RoleMiddleware(userentity.AdminRole),
	)
}
