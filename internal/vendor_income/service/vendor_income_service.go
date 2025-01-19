package vendorincomeservice

import (
	transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"
)

type IVendorIncomeService interface {
	CreateVendorIncome(transaction *transactionentity.Transaction) error
}
