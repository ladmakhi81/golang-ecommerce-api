package events

import transactionentity "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/entity"

type CalculateVendorIncomeEventBody struct {
	Transaction *transactionentity.Transaction
}

func NewCalculateVendorIncomeEventBody(
	transaction *transactionentity.Transaction,
) CalculateVendorIncomeEventBody {
	return CalculateVendorIncomeEventBody{
		Transaction: transaction,
	}
}

type PayedOrderEventBody struct {
	OrderId uint
}

func NewPayedOrderEventBody(orderId uint) PayedOrderEventBody {
	return PayedOrderEventBody{OrderId: orderId}
}
