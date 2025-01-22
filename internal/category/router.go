package category

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type CategoryRouter struct {
	baseApi    *echo.Group
	handler    CategoryHandler
	middleware middlewares.Middleware
}

func NewCategoryRouter(
	handler CategoryHandler,
	middleware middlewares.Middleware,
) CategoryRouter {
	return CategoryRouter{
		handler:    handler,
		middleware: middleware,
	}
}

func (userRouter *CategoryRouter) SetBaseApi(baseApi *echo.Group) {
	userRouter.baseApi = baseApi
}

func (categoryRouter CategoryRouter) RegisterRoutes() {
	categoriesApi := categoryRouter.baseApi.Group("/categories")

	categoriesApi.Use(
		categoryRouter.middleware.AuthMiddleware(),
	)

	categoriesApi.POST(
		"",
		categoryRouter.handler.CreateCategory,
		categoryRouter.middleware.RoleMiddleware(userentity.VendorRole),
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
		categoryRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
	categoriesApi.PATCH(
		"/icon/:categoryId",
		categoryRouter.handler.UploadCategoryIcon,
		categoryRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
}
