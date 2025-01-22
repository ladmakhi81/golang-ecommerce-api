package order

import (
	"fmt"

	"github.com/labstack/echo/v4"
	orderrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/order/repository"
	orderservice "github.com/ladmakhi81/golang-ecommerce-api/internal/order/service"
	"go.uber.org/dig"
)

type OrderModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewOrderModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) OrderModule {
	return OrderModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (orderModule OrderModule) LoadModule() {
	orderModule.diContainer.Provide(NewOrderRouter)
	orderModule.diContainer.Provide(NewOrderHandler)
	orderModule.diContainer.Provide(orderservice.NewOrderService)
	orderModule.diContainer.Provide(orderrepository.NewOrderRepository)
}

func (orderModule OrderModule) Run() {
	err := orderModule.diContainer.Invoke(func(orderRouter OrderRouter) {
		orderRouter.SetBaseApi(orderModule.baseApi)
		orderRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("OrderModule Loaded Successfully")
	} else {
		fmt.Println("OrderModule Not Load", err)
	}
}
