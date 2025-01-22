package vendorincomerepository

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/common/storage"
	vendorincomeentity "github.com/ladmakhi81/golang-ecommerce-api/internal/vendor_income/entity"
)

type VendorIncomeRepository struct {
	storage *storage.Storage
}

func NewVendorIncomeRepository(
	storage *storage.Storage,
) IVendorIncomeRepository {
	return VendorIncomeRepository{
		storage: storage,
	}
}

func (vendorIncomeRepo VendorIncomeRepository) CreateIncome(vendorIncome *vendorincomeentity.VendorIncome) error {
	command := `
		INSERT INTO _vendor_incomes (customer_id, order_amount, fee_amount, income_amount, order_item, transaction_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`
	row := vendorIncomeRepo.storage.DB.QueryRow(
		command,
		vendorIncome.Customer.ID,
		vendorIncome.OrderAmount,
		vendorIncome.FeeAmount,
		vendorIncome.IncomeAmount,
		vendorIncome.OrderItem.ID,
		vendorIncome.Transaction.ID,
	)
	return row.Err()
}
