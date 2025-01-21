package category

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type CategoryRouter struct {
	baseApi *echo.Group
	handler CategoryHandler
	config  config.MainConfig
}

func NewCategoryRouter(
	handler CategoryHandler,
	config config.MainConfig,
) CategoryRouter {
	return CategoryRouter{
		handler: handler,
		config:  config,
	}
}

func (userRouter *CategoryRouter) SetBaseApi(baseApi *echo.Group) {
	userRouter.baseApi = baseApi
}

func (categoryRouter CategoryRouter) RegisterRoutes() {
	categoriesApi := categoryRouter.baseApi.Group("/categories")

	categoriesApi.Use(
		middlewares.AuthMiddleware(
			categoryRouter.config.SecretKey,
		),
	)

	categoriesApi.POST(
		"",
		categoryRouter.handler.CreateCategory,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	categoriesApi.GET(
		"",
		categoryRouter.handler.GetCategoriesTree,
	)
	categoriesApi.GET(
		"/page",
		categoryRouter.handler.GetCategoriesPage,
	)
	categoriesApi.DELETE(
		"/:id",
		categoryRouter.handler.DeleteCategoryById,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	categoriesApi.PATCH(
		"/icon/:categoryId",
		categoryRouter.handler.UploadCategoryIcon,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
}
