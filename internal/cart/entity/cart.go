package cartentity

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	product_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Cart struct {
	Product   *product_entity.Product      `json:"product,omitempty"`
	Customer  *user_entity.User            `json:"customer,omitempty"`
	Quantity  uint                         `json:"quantity,omitempty"`
	PriceItem *product_entity.ProductPrice `json:"priceItem,omitempty"`

	entity.BaseEntity
}

func NewCart(
	product *product_entity.Product,
	customer *user_entity.User,
	priceItem *product_entity.ProductPrice,
	quantity uint,
) *Cart {
	return &Cart{
		Product:   product,
		Customer:  customer,
		Quantity:  quantity,
		PriceItem: priceItem,
	}
}
