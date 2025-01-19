package transactionevent

import (
	"fmt"

	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
)

type TransactionEventsSubscriber struct {
	transactionService transactionservice.ITransactionService
}

func NewTransactionEventsSubscriber(
	transactionService transactionservice.ITransactionService,
) TransactionEventsSubscriber {
	return TransactionEventsSubscriber{
		transactionService: transactionService,
	}
}

func (subscriber TransactionEventsSubscriber) SubscribeCalculateVendorIncome(event events.Event) {
	payload := event.Payload.(CalculateVendorIncomeEventBody)
	fmt.Println(payload)
}
