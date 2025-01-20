package category

import (
	"github.com/labstack/echo/v4"
	categoryservice "github.com/ladmakhi81/golang-ecommerce-api/internal/category/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
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
	categoriesApi := categoryRouter.apiRouter.Group("/categories")

	categoriesApi.Use(
		middlewares.AuthMiddleware(
			categoryRouter.config.SecretKey,
		),
	)

	categoriesApi.POST(
		"",
		categoryRouter.categoryHandler.CreateCategory,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	categoriesApi.GET(
		"",
		categoryRouter.categoryHandler.GetCategoriesTree,
	)
	categoriesApi.GET(
		"/page",
		categoryRouter.categoryHandler.GetCategoriesPage,
	)
	categoriesApi.DELETE(
		"/:id",
		categoryRouter.categoryHandler.DeleteCategoryById,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	categoriesApi.PATCH(
		"/icon/:categoryId",
		categoryRouter.categoryHandler.UploadCategoryIcon,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
}
