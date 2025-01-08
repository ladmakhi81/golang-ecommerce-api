package payment_entity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	order_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Payment struct {
	Customer        *user_entity.User
	Status          PaymentStatus
	StatusChangedAt time.Time
	Order           *order_entity.Order
	Amount          float32

	entity.BaseEntity
}
