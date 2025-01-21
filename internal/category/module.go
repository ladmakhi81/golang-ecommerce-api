package category

import (
	"fmt"

	"github.com/labstack/echo/v4"
	categoryrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/category/repository"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"go.uber.org/dig"
)

type CategoryModule struct {
	diContainer *dig.Container
	baseApi     *echo.Group
}

func NewCategoryModule(
	diContainer *dig.Container,
	baseApi *echo.Group,
) CategoryModule {
	return CategoryModule{
		diContainer: diContainer,
		baseApi:     baseApi,
	}
}

func (categoryModule CategoryModule) LoadModule() {
	categoryModule.diContainer.Provide(NewCategoryRouter)
	categoryModule.diContainer.Provide(NewCategoryHandler)
	categoryModule.diContainer.Provide(categoryservice.NewCategoryService)
	categoryModule.diContainer.Provide(categoryrepository.NewCategoryRepository)
}

func (categoryModule CategoryModule) Run() {
	err := categoryModule.diContainer.Invoke(func(categoryRouter CategoryRouter) {
		categoryRouter.SetBaseApi(categoryModule.baseApi)
		categoryRouter.RegisterRoutes()
	})

	if err == nil {
		fmt.Println("CategoryModule Loaded Successfully")
	} else {
		fmt.Println("CategoryModule Not Loaded", err)
	}
}
