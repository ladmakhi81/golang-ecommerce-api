package orderentity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Order struct {
	Customer        *user_entity.User
	Status          OrderStatus
	FinalPrice      float32
	StatusChangedAt time.Time
	Items           []*OrderItem

	entity.BaseEntity
}

func NewOrder(
	customer *user_entity.User,
	finalPrice float32,
) *Order {
	return &Order{
		Customer:   customer,
		FinalPrice: finalPrice,
	}
}
