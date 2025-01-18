package orderentity

import (
	product_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type OrderItem struct {
	ID        uint
	Product   *product_entity.Product
	PriceItem *product_entity.ProductPrice
	Vendor    *user_entity.User
	Customer  *user_entity.User
	Order     *Order
	Quantity  uint
}

func NewOrderItem(
	product *product_entity.Product,
	priceItem *product_entity.ProductPrice,
	vendor *user_entity.User,
	customer *user_entity.User,
	order *Order,
	quantity uint,
) *OrderItem {
	return &OrderItem{
		Product:   product,
		PriceItem: priceItem,
		Vendor:    vendor,
		Customer:  customer,
		Order:     order,
		Quantity:  quantity,
	}
}
