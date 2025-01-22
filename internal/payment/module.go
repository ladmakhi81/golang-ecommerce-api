package payment

import (
	"fmt"

	"github.com/labstack/echo/v4"
	paymentrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/repository"
	paymentservice "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/service"
	pkgzarinpalservice "github.com/ladmakhi81/golang-ecommerce-api/pkg/zarinpal/service"
	"go.uber.org/dig"
)

type PaymentModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewPaymentModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) PaymentModule {
	return PaymentModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (paymentModule PaymentModule) LoadModule() {
	paymentModule.diContainer.Provide(NewPaymentRouter)
	paymentModule.diContainer.Provide(NewPaymentHandler)
	paymentModule.diContainer.Provide(paymentservice.NewPaymentService)
	paymentModule.diContainer.Provide(paymentrepository.NewPaymentRepository)
	paymentModule.diContainer.Provide(pkgzarinpalservice.NewZarinpalService)
}

func (paymentModule PaymentModule) Run() {
	err := paymentModule.diContainer.Invoke(func(paymentRouter PaymentRouter) {
		paymentRouter.SetBaseApi(paymentModule.baseApi)
		paymentRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("PaymentModule Loaded Successfully")
	} else {
		fmt.Println("PaymentModule Not Load", err)
	}
}
