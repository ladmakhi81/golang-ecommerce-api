package vendorincomeentity

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/entity"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type VendorIncome struct {
	Customer     *userentity.User               `json:"customer,omitempty"`
	OrderAmount  float32                        `json:"orderAmount,omitempty"`
	FeeAmount    float32                        `json:"feeAmount,omitempty"`
	IncomeAmount float32                        `json:"incomeAmount,omitempty"`
	OrderItem    *orderentity.OrderItem         `json:"orderItem,omitempty"`
	Transaction  *transactionentity.Transaction `json:"transaction,omitempty"`

	entity.BaseEntity
}

func NewVendorIncome(
	customer *userentity.User,
	orderAmount float32,
	feeAmount float32,
	incomeAmount float32,
	orderItem *orderentity.OrderItem,
	transaction *transactionentity.Transaction,
) *VendorIncome {
	return &VendorIncome{
		Customer:     customer,
		OrderAmount:  orderAmount,
		FeeAmount:    feeAmount,
		IncomeAmount: incomeAmount,
		OrderItem:    orderItem,
		Transaction:  transaction,
	}
}
