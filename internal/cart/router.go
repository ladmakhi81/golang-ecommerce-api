package cart

import (
	"github.com/labstack/echo/v4"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/middlewares"
)

type CartRouter struct {
	baseApi    *echo.Group
	handler    CartHandler
	middleware middlewares.Middleware
}

func NewCartRouter(
	handler CartHandler,
	middleware middlewares.Middleware,
) CartRouter {
	return CartRouter{
		handler:    handler,
		middleware: middleware,
	}
}

func (cartRouter *CartRouter) SetBaseApi(baseApi *echo.Group) {
	cartRouter.baseApi = baseApi
}

func (cartRouter CartRouter) RegisterRoutes() {
	cartApi := cartRouter.baseApi.Group("/cart")

	cartApi.Use(
		cartRouter.middleware.AuthMiddleware(),
	)
	cartApi.POST("", cartRouter.handler.AddProductToCart)
	cartApi.DELETE("/:cartId", cartRouter.handler.DeleteUserCart)
	cartApi.PATCH("/:cartId", cartRouter.handler.UpdateCartQuantity)
	cartApi.GET("", cartRouter.handler.GetUserCarts)
}
