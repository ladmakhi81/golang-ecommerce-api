package transactionrepository

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
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
		INSERT INTO _transactions(user_id, payment_id, order_id, authority, ref_id, amount)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`
	row := transactionRepo.storage.DB.QueryRow(
		command,
		transaction.Customer.ID,
		transaction.Payment.ID,
		transaction.Order.ID,
		transaction.Authority,
		transaction.RefID,
		transaction.Amount,
	)
	scanErr := row.Scan(
		&transaction.ID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	return scanErr
}
func (transactionRepo TransactionRepository) GetTransactionsPage(page, limit uint) ([]*transactionentity.Transaction, error) {
	command := `
		SELECT 
		t.id, t.created_at, t.updated_at, t.authority, t.amount, t.ref_id,
		u.id, u.email, u.created_at, u.updated_at,
		p.id, p.merchant_id, p.created_at, p.updated_at,
		o.id, o.created_at, o.updated_at
		FROM _transactions t
		INNER JOIN _users u ON u.id = t.user_id
		INNER JOIN _payments p ON p.id = t.payment_id
		INNER JOIN _orders o ON o.id = t.order_id
		LIMIT $1 OFFSET $2
	`
	rows, err := transactionRepo.storage.DB.Query(command, limit, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := []*transactionentity.Transaction{}
	for rows.Next() {
		transaction := new(transactionentity.Transaction)
		transaction.Customer = new(userentity.User)
		transaction.Payment = new(paymententity.Payment)
		transaction.Order = new(orderentity.Order)
		scanErr := rows.Scan(
			&transaction.ID, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.Authority, &transaction.Amount, &transaction.RefID,
			&transaction.Customer.ID, &transaction.Customer.Email, &transaction.Customer.CreatedAt, &transaction.Customer.UpdatedAt,
			&transaction.Payment.ID, &transaction.Payment.MerchantID, &transaction.Payment.CreatedAt, &transaction.Payment.UpdatedAt,
			&transaction.Order.ID, &transaction.Order.CreatedAt, &transaction.Order.UpdatedAt,
		)

		if scanErr != nil {
			return nil, scanErr
		}

		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
