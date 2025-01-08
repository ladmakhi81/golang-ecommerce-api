package order

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/product"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
)

type Order struct {
	Customer        *user.User
	Status          OrderStatus
	FinalPrice      float32
	StatusChangedAt time.Time
	Items           []*OrderItem

	entity.BaseEntity
}

type OrderItem struct {
	ID       uint
	Product  *product.Product
	Vendor   *user.User
	Quantity uint
}

type OrderStatus string

const (
	OrderStatusPending     = "Pending"
	OrderStatusPayed       = "Payed"
	OrderStatusPreparation = "Preparation"
	OrderStatusDelivery    = "Delivery"
	OrderStatusDone        = "Done"
)

func IsValid(status OrderStatus) bool {
	validStatuses := []OrderStatus{
		OrderStatusDelivery,
		OrderStatusDone,
		OrderStatusPayed,
		OrderStatusPending,
		OrderStatusPreparation,
	}

	for _, validStatus := range validStatuses {
		if validStatus == status {
			return true
		}
	}

	return false
}
