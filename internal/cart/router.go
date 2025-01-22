package cart

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
)

type CartRouter struct {
	baseApi *echo.Group
	handler CartHandler
	config  config.MainConfig
}

func NewCartRouter(
	handler CartHandler,
	config config.MainConfig,
) CartRouter {
	return CartRouter{
		config:  config,
		handler: handler,
	}
}

func (cartRouter *CartRouter) SetBaseApi(baseApi *echo.Group) {
	cartRouter.baseApi = baseApi
}

func (cartRouter CartRouter) RegisterRoutes() {
	cartApi := cartRouter.baseApi.Group("/cart")

	cartApi.Use(
		middlewares.AuthMiddleware(
			cartRouter.config.SecretKey,
		),
	)
	cartApi.POST("", cartRouter.handler.AddProductToCart)
	cartApi.DELETE("/:cartId", cartRouter.handler.DeleteUserCart)
	cartApi.PATCH("/:cartId", cartRouter.handler.UpdateCartQuantity)
	cartApi.GET("", cartRouter.handler.GetUserCarts)
}
