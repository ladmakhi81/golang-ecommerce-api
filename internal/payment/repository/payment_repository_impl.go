package paymentrepository

import (
	"database/sql"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	orderentity "github.com/ladmakhi81/golang-ecommerce-api/internal/order/entity"
	paymententity "github.com/ladmakhi81/golang-ecommerce-api/internal/payment/entity"
	userentity "github.com/ladmakhi81/golang-ecommerce-api/internal/user/entity"
)

type PaymentRepository struct {
	storage *storage.Storage
}

func NewPaymentRepository(storage *storage.Storage) PaymentRepository {
	return PaymentRepository{
		storage: storage,
	}
}

func (paymentRepo PaymentRepository) CreatePayment(payment *paymententity.Payment) error {
	command := `
		INSERT INTO _payments (customer_id, status, order_id, amount, authority, merchant_id) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`
	row := paymentRepo.storage.DB.QueryRow(
		command,
		payment.Customer.ID,
		payment.Status,
		payment.Order.ID,
		payment.Amount,
		payment.Authority,
		payment.MerchantID,
	)
	scanErr := row.Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt)
	return scanErr
}
func (paymentRepo PaymentRepository) FindPaymentByAuthority(authority string) (*paymententity.Payment, error) {
	command := `
		SELECT
		p.id, p.created_at, p.updated_at, p.status, p.status_changed_at, p.amount, p.authority, p.merchant_id,
		u.id, u.email, u.created_at, u.updated_at,
		o.id, o.final_price, o.created_at, o.updated_at
		FROM _payments p
		INNER JOIN _users u ON u.id = p.customer_id
		INNER JOIN _orders o ON o.id = p.order_id
		WHERE authority = $1
		LIMIT 1
	`
	row := paymentRepo.storage.DB.QueryRow(command, authority)
	payment := new(paymententity.Payment)
	payment.Customer = new(userentity.User)
	payment.Order = new(orderentity.Order)
	scanErr := row.Scan(
		&payment.ID,
		&payment.CreatedAt,
		&payment.UpdatedAt,
		&payment.Status,
		&payment.StatusChangedAt,
		&payment.Amount,
		&payment.Authority,
		&payment.MerchantID,
		&payment.Customer.ID,
		&payment.Customer.Email,
		&payment.Customer.CreatedAt,
		&payment.Customer.UpdatedAt,
		&payment.Order.ID,
		&payment.Order.FinalPrice,
		&payment.Order.CreatedAt,
		&payment.Order.UpdatedAt,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	return payment, nil
}
func (paymentRepo PaymentRepository) UpdatePaymentStatus(payment *paymententity.Payment) error {
	command := `
		UPDATE _payments SET
		status = $1, status_changed_at = $2
		WHERE id = $3
	`
	row := paymentRepo.storage.DB.QueryRow(command, payment.Status, payment.StatusChangedAt, payment.ID)
	return row.Err()
}
