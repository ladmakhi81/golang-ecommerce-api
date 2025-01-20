package transactionservice

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	transactionrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/repository"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type TransactionService struct {
	transactionRepo transactionrepository.ITransactionRepository
}

func NewTransactionService(
	transactionRepo transactionrepository.ITransactionRepository,
) TransactionService {
	return TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (transactionService TransactionService) CreateTransaction(payment *paymententity.Payment, refId uint, user *userentity.User) (*transactionentity.Transaction, error) {
	transaction := transactionentity.NewTransaction(
		user,
		payment,
		payment.Order,
		payment.Authority,
		refId,
		payment.Amount,
	)
	if transactionErr := transactionService.transactionRepo.CreateTransaction(transaction); transactionErr != nil {
		return nil, types.NewServerError(
			"error in creating transaction",
			"TransactionService.CreateTransaction",
			transactionErr,
		)
	}
	return transaction, nil
}
func (transactionService TransactionService) GetTransactionsPage(page, limit uint) ([]*transactionentity.Transaction, uint, error) {
	transactions, transactionsErr := transactionService.transactionRepo.GetTransactionsPage(page, limit)
	if transactionsErr != nil {
		return nil, 0, types.NewServerError(
			"error in finding transactions page",
			"TransactionService.GetTransactionsPage",
			transactionsErr,
		)
	}
	transactionCount, transactionCountErr := transactionService.transactionRepo.GetTransactionsCount()
	if transactionCountErr != nil {
		return nil, 0, types.NewServerError(
			"error in get count of transactions",
			"TransactionService.GetTransactionsPage.GetTransactionCount",
			transactionCountErr,
		)
	}
	return transactions, transactionCount, nil
}
func (transactionService TransactionService) GetOrderIdOfTransaction(transactionId uint) (*uint, error) {
	orderId, orderIdErr := transactionService.transactionRepo.GetOrderIdOfTransaction(transactionId)
	if orderIdErr != nil {
		return nil, types.NewServerError(
			"error in finding order id based on transaction id",
			"TransactionService.GetOrderIdOfTransaction",
			orderIdErr,
		)
	}
	return orderId, nil
}
