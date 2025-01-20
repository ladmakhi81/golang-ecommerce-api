package orderentity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Order struct {
	Customer        *user_entity.User        `json:"customer,omitempty"`
	Status          OrderStatus              `json:"status,omitempty"`
	FinalPrice      float32                  `json:"finalPrice,omitempty"`
	StatusChangedAt time.Time                `json:"statusChangedAt,omitempty"`
	Items           []*OrderItem             `json:"items,omitempty"`
	Address         *user_entity.UserAddress `json:"address,omitempty"`

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
