package transactionservice

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/types"
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	transactionrepository "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/repository"
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

func (transactionService TransactionService) CreatePaymentTransaction(payment *paymententity.Payment, refId uint) (*transactionentity.Transaction, error) {
	transaction := transactionentity.NewTransaction(
		payment.Customer,
		payment,
		payment.Order,
		payment.Authority,
		refId,
		payment.Amount,
		transactionentity.TransactionTypePayment,
	)
	if transactionErr := transactionService.transactionRepo.CreateTransaction(transaction); transactionErr != nil {
		return nil, types.NewServerError(
			"error in creating transaction",
			"TransactionService.CreatePaymentTransaction.CreateTransaction",
			transactionErr,
		)
	}
	return transaction, nil
}
