package order_entity

import (
	product_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type OrderItem struct {
	ID       uint
	Product  *product_entity.Product
	Vendor   *user_entity.User
	Quantity uint
}
