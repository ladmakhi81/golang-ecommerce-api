package product

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
	productservice "github.com/ladmakhi81/golang-ecommerce-api/internal/product/service"
)

type ProductRouter struct {
	apiRouter      *echo.Group
	productHandler ProductHandler
	productService productservice.IProductService
	util           utils.Util
	config         config.MainConfig
}

func NewProductRouter(
	apiRouter *echo.Group,
	config config.MainConfig,
	productService productservice.IProductService,
) ProductRouter {
	return ProductRouter{
		apiRouter:      apiRouter,
		productHandler: NewProductHandler(productService),
		util:           utils.NewUtil(),
		config:         config,
	}
}

func (productRouter ProductRouter) SetupRouter() {
	productsApi := productRouter.apiRouter.Group("/products")

	productsApi.POST(
		"",
		productRouter.productHandler.CreateProduct,
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)

	productsApi.PATCH(
		"/:id",
		productRouter.productHandler.ConfirmProductByAdmin,
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)

	productsApi.GET(
		"/:id",
		productRouter.productHandler.FindProductDetailById,
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)

	productsApi.GET(
		"",
		productRouter.productHandler.GetProductsPage,
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)

	productsApi.DELETE(
		"/:id",
		productRouter.productHandler.DeleteProductById,
		middlewares.AuthMiddleware(
			productRouter.config.SecretKey,
		),
	)
}
