package cartrepository

import cartentity "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/entity"

type ICartRepository interface {
	CreateProductCart(cart *cartentity.Cart) error
	FindCustomerCartByProductIdAndPriceId(customerID, productID uint, priceId uint) (*cartentity.Cart, error)
	UpdateCartById(cart *cartentity.Cart) error
	DeleteCartById(cartID uint) error
	FindCartById(cartID uint) (*cartentity.Cart, error)
	FindCustomerCart(customerId uint) ([]*cartentity.Cart, error)
	FindCartsByIds(ids []uint) ([]*cartentity.Cart, error)
	DeleteCartsByIds(ids []uint) error
}
