package order

import (
	"fmt"

	"github.com/labstack/echo/v4"
	orderevent "github.com/ladmakhi81/golang-ecommerce-api/internal/order/event"
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
	orderModule.diContainer.Provide(orderevent.NewOrderEventsSubscriber)
	orderModule.diContainer.Provide(orderevent.NewOrderEventsContainer)
}

func (orderModule OrderModule) Run() {
	err := orderModule.diContainer.Invoke(func(orderRouter OrderRouter) {
		orderRouter.SetBaseApi(orderModule.baseApi)
		orderRouter.RegisterRoutes()
	})

	orderEventErr := orderModule.diContainer.Invoke(func(orderEventsContainer orderevent.OrderEventsContainer) {
		orderEventsContainer.RegisterEvents()
	})

	if err == nil && orderEventErr == nil {
		fmt.Println("OrderModule Loaded Successfully")
	} else {
		fmt.Println("OrderModule Not Load")
	}
}
