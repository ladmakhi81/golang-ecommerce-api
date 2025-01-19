package paymententity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	order_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Payment struct {
	Customer        *user_entity.User   `json:"customer,omitempty"`
	Status          PaymentStatus       `json:"status,omitempty"`
	StatusChangedAt time.Time           `json:"statusChangedAt,omitempty"`
	Order           *order_entity.Order `json:"order,omitempty"`
	Amount          float32             `json:"amount,omitempty"`
	Authority       string              `json:"authority,omitempty"`
	MerchantID      string              `json:"merchant_id,omitempty"`

	entity.BaseEntity
}

func NewPayment(
	order *order_entity.Order,
	authority,
	merchantID string,
) *Payment {
	return &Payment{
		Customer:   order.Customer,
		Status:     PaymentStatusPending,
		Order:      order,
		Amount:     order.FinalPrice,
		Authority:  authority,
		MerchantID: merchantID,
	}
}
