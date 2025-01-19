package payment

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
)

type PaymentRouter struct {
	apiRouter      *echo.Group
	paymentHandler PaymentHandler
	config         config.MainConfig
	paymentService paymentservice.IPaymentService
}

func NewPaymentRouter(
	apiRouter *echo.Group,
	config config.MainConfig,
	paymentService paymentservice.IPaymentService,
) PaymentRouter {
	return PaymentRouter{
		apiRouter:      apiRouter,
		config:         config,
		paymentHandler: NewPaymentHandler(paymentService),
	}
}

func (paymentRouter PaymentRouter) SetupRouter() {
	paymentsApi := paymentRouter.apiRouter.Group("/payments")

	paymentsApi.Use(
		middlewares.AuthMiddleware(paymentRouter.config.SecretKey),
	)

	paymentsApi.POST(
		"/verify",
		paymentRouter.paymentHandler.VerifyPayment,
	)
	paymentsApi.GET(
		"/page",
		paymentRouter.paymentHandler.GetPaymentsPage,
	)
}
