package transactionrepository

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
)

type TransactionRepository struct {
	storage *storage.Storage
}

func NewTransactionRepository(
	storage *storage.Storage,
) TransactionRepository {
	return TransactionRepository{
		storage: storage,
	}
}

func (transactionRepo TransactionRepository) CreateTransaction(transaction *transactionentity.Transaction) error {
	command := `
		INSERT INTO _transactions(user_id, payment_id, order_id, authority, ref_id, amount, type)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at;
	`
	row := transactionRepo.storage.DB.QueryRow(
		command,
		transaction.User.ID,
		transaction.Payment.ID,
		transaction.Order.ID,
		transaction.Authority,
		transaction.RefID,
		transaction.Amount,
		transaction.Type,
	)
	scanErr := row.Scan(
		&transaction.ID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	return scanErr
}
