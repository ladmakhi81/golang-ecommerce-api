package cart

import (
	"github.com/labstack/echo/v4"
	cartservice "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/service"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/config"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
)

type CartRouter struct {
	apiRouter   *echo.Group
	cartHandler CartHandler
	config      config.MainConfig
	cartService cartservice.ICartService
}

func NewCartRouter(
	apiRouter *echo.Group,
	config config.MainConfig,
	cartService cartservice.ICartService,
) CartRouter {
	return CartRouter{
		apiRouter:   apiRouter,
		config:      config,
		cartHandler: NewCartHandler(cartService),
	}
}

func (cartRouter CartRouter) Setup() {
	cartApi := cartRouter.apiRouter.Group("/cart")
	cartApi.Use(
		middlewares.AuthMiddleware(
			cartRouter.config.SecretKey,
		),
	)
	cartApi.POST("", cartRouter.cartHandler.AddProductToCart)
	cartApi.DELETE("/:cartId", cartRouter.cartHandler.DeleteUserCart)
	cartApi.PATCH("/:cartId", cartRouter.cartHandler.UpdateCartQuantity)
	cartApi.GET("", cartRouter.cartHandler.GetUserCarts)
}
