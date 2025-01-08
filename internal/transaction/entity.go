package transaction

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/order"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/payment"
	"github.com/ladmakhi81/golang-ecommerce-api/internal/user"
)

type Transaction struct {
	Payer     *user.User
	Receiver  *user.User
	Payment   *payment.Payment
	Order     *order.Order
	Authority string
	RefID     string
	Amount    float32

	entity.BaseEntity
}
