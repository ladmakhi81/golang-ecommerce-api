package transactionrepository

import transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"

type ITransactionRepository interface {
	CreateTransaction(transaction *transactionentity.Transaction) error
	GetTransactionsPage(page, limit uint) ([]*transactionentity.Transaction, error)
	GetOrderIdOfTransaction(transactionId uint) (*uint, error)
	GetTransactionsCount() (uint, error)
}
