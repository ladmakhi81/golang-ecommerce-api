package product

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type ProductRouter struct {
	baseApi *echo.Group
	handler ProductHandler
	util    utils.Util
	config  config.MainConfig
}

func NewProductRouter(
	config config.MainConfig,
	handler ProductHandler,
	util utils.Util,
) ProductRouter {
	return ProductRouter{
		handler: handler,
		util:    util,
		config:  config,
	}
}

func (productRouter *ProductRouter) SetBaseApi(baseApi *echo.Group) {
	productRouter.baseApi = baseApi
}

func (productRouter ProductRouter) RegisterRoutes() {
	productsApi := productRouter.baseApi.Group("/products")

	productsApi.Use(
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)
	productsApi.POST(
		"",
		productRouter.handler.CreateProduct,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.PATCH(
		"/:id",
		productRouter.handler.ConfirmProductByAdmin,
		middlewares.RoleMiddleware(userentity.AdminRole),
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
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.POST(
		"/price/:product_id",
		productRouter.handler.AddPriceToProductPriceList,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.DELETE(
		"/price/:id",
		productRouter.handler.DeletePriceItem,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
	productsApi.PATCH(
		"/images/:id",
		productRouter.handler.UploadProductImages,
		middlewares.RoleMiddleware(userentity.VendorRole),
	)
}
