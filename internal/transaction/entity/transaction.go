package transaction_entity

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	order_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	payment_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Transaction struct {
	Payer     *user_entity.User
	Receiver  *user_entity.User
	Payment   *payment_entity.Payment
	Order     *order_entity.Order
	Authority string
	RefID     string
	Amount    float32

	entity.BaseEntity
}
