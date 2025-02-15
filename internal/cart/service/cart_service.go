package cartservice

import (
	cartdto "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/dto"
	cartentity "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/entity"
)

type ICartService interface {
	AddProductToCart(customerID uint, reqBody cartdto.AddProductCartReqBody) (*cartentity.Cart, error)
	FindCustomerCartByProductIdAndPriceId(customerID, productID uint, priceID uint) (*cartentity.Cart, error)
	UpdateCartQuantityById(customerID, cartId uint, reqBody cartdto.UpdateCartQuantityReqBody) error
	DeleteCartById(customerID, cartId uint) error
	FindCartById(cartId uint) (*cartentity.Cart, error)
	FindCustomerCart(customerId uint) ([]*cartentity.Cart, error)
	FindCartsByIds(ids []uint) ([]*cartentity.Cart, error)
	CalculateFinalPriceOfCarts(carts []*cartentity.Cart) float32
	DeleteCartsByIds(ids []uint) error
}
