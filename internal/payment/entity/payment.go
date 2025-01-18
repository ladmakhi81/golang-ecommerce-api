package paymententity

import (
	"time"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	order_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Payment struct {
	Customer        *user_entity.User   `json:"customer"`
	Status          PaymentStatus       `json:"status"`
	StatusChangedAt time.Time           `json:"statusChangedAt"`
	Order           *order_entity.Order `json:"order"`
	Amount          float32             `json:"amount"`
	Authority       string              `json:"authority"`
	MerchantID      string              `json:"merchant_id"`

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
