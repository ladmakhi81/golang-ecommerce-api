package product_entity

import "github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"

type ProductPrice struct {
	Key        string
	Value      string
	ExtraPrice float32

	entity.BaseEntity
}
