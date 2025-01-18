package transactionservice

import (
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
)

type ITransactionService interface {
	CreatePaymentTransaction(payment *paymententity.Payment, refId uint) (*transactionentity.Transaction, error)
}
