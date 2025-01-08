package cart_entity

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	product_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/product/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Cart struct {
	Product  product_entity.Product
	Customer user_entity.User
	Quantity uint

	entity.BaseEntity
}
