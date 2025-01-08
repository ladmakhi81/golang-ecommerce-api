package payment

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/order"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
)

type Payment struct {
	Customer        *user.User
	Status          PaymentStatus
	StatusChangedAt time.Time
	Order           *order.Order
	Amount          float32

	entity.BaseEntity
}

type PaymentStatus string

const (
	PaymentStatusPending = "Pending"
	PaymentStatusSuccess = "Success"
	PaymentStatusFailed  = "Failed"
)

func IsValid(status PaymentStatus) bool {
	validStatuses := []PaymentStatus{
		PaymentStatusFailed,
		PaymentStatusPending,
		PaymentStatusSuccess,
	}

	for _, validStatus := range validStatuses {
		if validStatus == status {
			return true
		}
	}

	return false
}
