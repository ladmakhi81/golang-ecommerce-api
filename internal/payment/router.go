package payment

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type PaymentRouter struct {
	baseApi    *echo.Group
	handler    PaymentHandler
	middleware middlewares.Middleware
}

func NewPaymentRouter(
	handler PaymentHandler,
	middleware middlewares.Middleware,
) PaymentRouter {
	return PaymentRouter{
		handler:    handler,
		middleware: middleware,
	}
}

func (paymentRouter *PaymentRouter) SetBaseApi(baseApi *echo.Group) {
	paymentRouter.baseApi = baseApi
}

func (paymentRouter PaymentRouter) RegisterRoutes() {
	paymentsApi := paymentRouter.baseApi.Group("/payments")

	paymentsApi.Use(
		paymentRouter.middleware.AuthMiddleware(),
	)

	paymentsApi.POST(
		"/verify",
		paymentRouter.handler.VerifyPayment,
	)
	paymentsApi.GET(
		"/page",
		paymentRouter.handler.GetPaymentsPage,
		paymentRouter.middleware.RoleMiddleware(userentity.AdminRole),
	)
}
