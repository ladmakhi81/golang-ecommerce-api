package cart

import (
	"net/http"

	"github.com/labstack/echo/v4"
	cartdto "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/dto"
	cartservice "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/service"
	responsehandling "github.com/ladmakhi81/golang-ecommerce-api/internal/common/response_handling"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/utils"
)

type CartHandler struct {
	cartService cartservice.ICartService
	util        utils.Util
}

func NewCartHandler(
	cartService cartservice.ICartService,
) CartHandler {
	return CartHandler{
		cartService: cartService,
		util:        utils.NewUtil(),
	}
}

func (cartHandler CartHandler) AddProductToCart(c echo.Context) error {
	var reqBody cartdto.AddProductCartReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	customerId := c.Get("AuthClaim").(*types.AuthClaim).ID
	cart, cartErr := cartHandler.cartService.AddProductToCart(customerId, reqBody)
	if cartErr != nil {
		return cartErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		cart,
	)
	return nil
}
func (cartHandler CartHandler) DeleteUserCart(c echo.Context) error {
	cartId, parseCartIdErr := cartHandler.util.NumericParamConvertor(
		c.Param("cartId"),
		"invalid cart id",
	)
	if parseCartIdErr != nil {
		return parseCartIdErr
	}
	customerId := c.Get("AuthClaim").(*types.AuthClaim).ID
	deleteErr := cartHandler.cartService.DeleteCartById(customerId, cartId)
	if deleteErr != nil {
		return deleteErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
func (cartHandler CartHandler) UpdateCartQuantity(c echo.Context) error {
	cartId, parseCartErr := cartHandler.util.NumericParamConvertor(
		c.Param("cartId"),
		"invalid cart id",
	)
	if parseCartErr != nil {
		return parseCartErr
	}
	var reqBody cartdto.UpdateCartQuantityReqBody
	if err := c.Bind(&reqBody); err != nil {
		return types.NewClientError("invalid request body", http.StatusBadRequest)
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	customerId := c.Get("AuthClaim").(*types.AuthClaim).ID
	if updateErr := cartHandler.cartService.UpdateCartQuantityById(customerId, cartId, reqBody); updateErr != nil {
		return updateErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		nil,
	)
	return nil
}
func (cartHandler CartHandler) GetUserCarts(c echo.Context) error {
	customerId := c.Get("AuthClaim").(*types.AuthClaim).ID
	carts, cartsErr := cartHandler.cartService.FindCustomerCart(customerId)
	if cartsErr != nil {
		return cartsErr
	}
	responsehandling.ResponseJSON(
		c,
		http.StatusOK,
		carts,
	)
	return nil
}
