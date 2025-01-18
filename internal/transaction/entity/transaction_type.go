package transactionentity

type TransactionType string

const (
	TransactionTypeVendorIncome TransactionType = "VendorIncome"
	TransactionTypePayment      TransactionType = "Payment"
)

func (TransactionType) IsValid(value TransactionType) bool {
	return !(value == "VendorIncome" || value == "Payment")
}
