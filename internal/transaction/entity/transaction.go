package transactionentity

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	order_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	payment_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	user_entity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type Transaction struct {
	User      *user_entity.User       `json:"user"`
	Payment   *payment_entity.Payment `json:"payment"`
	Order     *order_entity.Order     `json:"order"`
	Authority string                  `json:"authority"`
	RefID     uint                    `json:"refId"`
	Amount    float32                 `json:"amount"`
	Type      TransactionType         `json:"type"`

	entity.BaseEntity
}

func NewTransaction(
	user *user_entity.User,
	payment *payment_entity.Payment,
	order *order_entity.Order,
	authority string,
	refID uint,
	amount float32,
	transactionType TransactionType,
) *Transaction {
	return &Transaction{
		User:      user,
		Payment:   payment,
		Order:     order,
		Authority: authority,
		RefID:     refID,
		Amount:    amount,
		Type:      transactionType,
	}
}
