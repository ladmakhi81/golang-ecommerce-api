package product

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
)

type ProductRouter struct {
	apiRouter           *echo.Group
	productHandler      ProductHandler
	productService      productservice.IProductService
	productPriceService productservice.IProductPriceService
	util                utils.Util
	config              config.MainConfig
}

func NewProductRouter(
	apiRouter *echo.Group,
	config config.MainConfig,
	productService productservice.IProductService,
	productPriceService productservice.IProductPriceService,
) ProductRouter {
	return ProductRouter{
		apiRouter:      apiRouter,
		productHandler: NewProductHandler(productService, productPriceService),
		util:           utils.NewUtil(),
		config:         config,
	}
}

func (productRouter ProductRouter) SetupRouter() {
	productsApi := productRouter.apiRouter.Group("/products")
	productsApi.Use(
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)

	productsApi.POST(
		"",
		productRouter.productHandler.CreateProduct,
	)
	productsApi.PATCH(
		"/:id",
		productRouter.productHandler.ConfirmProductByAdmin,
	)
	productsApi.GET(
		"/prices/:productId",
		productRouter.productHandler.GetPricesOfProduct,
	)
	productsApi.GET(
		"/:id",
		productRouter.productHandler.FindProductDetailById,
	)
	productsApi.GET(
		"",
		productRouter.productHandler.GetProductsPage,
	)
	productsApi.DELETE(
		"/:id",
		productRouter.productHandler.DeleteProductById,
	)
	productsApi.POST(
		"/price/:product_id",
		productRouter.productHandler.AddPriceToProductPriceList,
	)
	productsApi.DELETE(
		"/price/:id",
		productRouter.productHandler.DeletePriceItem,
	)
	productsApi.PATCH(
		"/images/:id",
		productRouter.productHandler.UploadProductImages,
	)
}
