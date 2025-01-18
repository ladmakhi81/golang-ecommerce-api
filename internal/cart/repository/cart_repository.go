package cartrepository

import cartentity "github.com/ladmakhi81/golang-ecommerce-api/internal/cart/entity"

type ICartRepository interface {
	CreateProductCart(cart *cartentity.Cart) error
	FindCustomerCartByProductId(customerID, productID uint) (*cartentity.Cart, error)
	UpdateCartById(cart *cartentity.Cart) error
	DeleteCartById(cartID uint) error
	FindCartById(cartID uint) (*cartentity.Cart, error)
	FindCustomerCart(customerId uint) ([]*cartentity.Cart, error)
}
