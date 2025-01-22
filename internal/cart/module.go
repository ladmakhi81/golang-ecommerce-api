package cart

import (
	"fmt"

	"github.com/labstack/echo/v4"
	cartrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/repository"
	cartservice "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/service"
	"go.uber.org/dig"
)

type CartModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewCartModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) CartModule {
	return CartModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (cartModule CartModule) LoadModule() {
	cartModule.diContainer.Provide(NewCartRouter)
	cartModule.diContainer.Provide(NewCartHandler)
	cartModule.diContainer.Provide(cartservice.NewCartService)
	cartModule.diContainer.Provide(cartrepository.NewCartRepository)
}

func (cartModule CartModule) Run() {
	err := cartModule.diContainer.Invoke(func(cartRouter CartRouter) {
		cartRouter.SetBaseApi(cartModule.baseApi)
		cartRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("CartModule Loaded Successfully")
	} else {
		fmt.Println("CartModule Not Load", err)
	}
}
