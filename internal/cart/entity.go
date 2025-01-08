package cart

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/product"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
)

type Cart struct {
	Product  product.Product
	Customer user.User
	Quantity uint

	entity.BaseEntity
}
