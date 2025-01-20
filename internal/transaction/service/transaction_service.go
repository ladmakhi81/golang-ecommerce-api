package transactionservice

import (
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type ITransactionService interface {
	CreateTransaction(payment *paymententity.Payment, refId uint, user *userentity.User) (*transactionentity.Transaction, error)
	GetTransactionsPage(page, limit uint) ([]*transactionentity.Transaction, uint, error)
	GetOrderIdOfTransaction(transactionId uint) (*uint, error)
}
