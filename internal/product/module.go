package product

import (
	"fmt"

	"github.com/labstack/echo/v4"
	productrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/product/repository"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
	"go.uber.org/dig"
)

type ProductModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewProductModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) ProductModule {
	return ProductModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (productModule ProductModule) LoadModule() {
	productModule.diContainer.Provide(productservice.NewProductService)
	productModule.diContainer.Provide(productservice.NewProductPriceService)
	productModule.diContainer.Provide(productrepository.NewProductPriceRepository)
	productModule.diContainer.Provide(productrepository.NewProductRepository)
	productModule.diContainer.Provide(NewProductHandler)
	productModule.diContainer.Provide(NewProductRouter)
}

func (productModule ProductModule) Run() {
	err := productModule.diContainer.Invoke(func(productRouter ProductRouter) {
		productRouter.SetBaseApi(productModule.baseApi)
		productRouter.RegisterRoutes()
	})
	if err == nil {
		fmt.Println("ProductModule Loaded Successfully")
	} else {
		fmt.Println("ProductModule Not Loaded", err)
	}
}
