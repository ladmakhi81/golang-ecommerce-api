package category

import (
	"github.com/labstack/echo/v4"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
)

type CategoryRouter struct {
	apiRouter       *echo.Group
	config          config.MainConfig
	categoryHandler CategoryHandler
}

func NewCategoryRouter(
	apiRouter *echo.Group,
	config config.MainConfig,
	categoryService categoryservice.ICategoryService,
) CategoryRouter {
	return CategoryRouter{
		apiRouter:       apiRouter,
		categoryHandler: NewCategoryHandler(categoryService),
		config:          config,
	}
}

func (categoryRouter CategoryRouter) SetupRouter() {
	categoriesApiRouter := categoryRouter.apiRouter.Group("/categories")

	categoriesApiRouter.POST(
		"",
		categoryRouter.categoryHandler.CreateCategory,
		middlewares.AuthMiddleware(
			categoryRouter.config.SecretKey,
		),
	)

	categoriesApiRouter.GET(
		"",
		categoryRouter.categoryHandler.GetCategoriesTree,
		middlewares.AuthMiddleware(
			categoryRouter.config.SecretKey,
		),
	)

	categoriesApiRouter.GET(
		"/page",
		categoryRouter.categoryHandler.GetCategoriesPage,
		middlewares.AuthMiddleware(
			categoryRouter.config.SecretKey,
		),
	)
	categoriesApiRouter.DELETE(
		"/:id",
		categoryRouter.categoryHandler.DeleteCategoryById,
		middlewares.AuthMiddleware(
			categoryRouter.config.SecretKey,
		),
	)
}
