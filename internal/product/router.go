package product

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type ProductRouter struct {
	baseApi    *echo.Group
	handler    ProductHandler
	util       utils.Util
	middleware middlewares.Middleware
}

func NewProductRouter(
	handler ProductHandler,
	util utils.Util,
	middleware middlewares.Middleware,
) ProductRouter {
	return ProductRouter{
		handler:    handler,
		util:       util,
		middleware: middleware,
	}
}

func (productRouter *ProductRouter) SetBaseApi(baseApi *echo.Group) {
	productRouter.baseApi = baseApi
}

func (productRouter ProductRouter) RegisterRoutes() {
	productsApi := productRouter.baseApi.Group("/products")

	productsApi.Use(
		productRouter.middleware.AuthMiddleware(),
	)
	productsApi.POST(
		"",
		productRouter.handler.CreateProduct,
		productRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.PATCH(
		"/:id",
		productRouter.handler.ConfirmProductByAdmin,
		productRouter.middleware.RoleMiddleware(userentity.AdminRole),
	)
	productsApi.GET(
		"/prices/:productId",
		productRouter.handler.GetPricesOfProduct,
	)
	productsApi.GET(
		"/:id",
		productRouter.handler.FindProductDetailById,
	)
	productsApi.GET(
		"",
		productRouter.handler.GetProductsPage,
	)
	productsApi.DELETE(
		"/:id",
		productRouter.handler.DeleteProductById,
		productRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.POST(
		"/price/:product_id",
		productRouter.handler.AddPriceToProductPriceList,
		productRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.DELETE(
		"/price/:id",
		productRouter.handler.DeletePriceItem,
		productRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.PATCH(
		"/images/:id",
		productRouter.handler.UploadProductImages,
		productRouter.middleware.RoleMiddleware(userentity.VendorRole),
	)
}
