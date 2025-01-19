package transactionevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
	transactionservice "github.com/ladmakhi81/golang-ecommerce-api/internal/transaction/service"
)

type TransactionEventsContainer struct {
	eventContainer              *events.EventsContainer
	transactionService          transactionservice.ITransactionService
	transactionEventsSubscriber TransactionEventsSubscriber
}

func NewTransactionEventsContainer(
	eventContainer *events.EventsContainer,
	transactionService transactionservice.ITransactionService,
) TransactionEventsContainer {
	return TransactionEventsContainer{
		eventContainer:              eventContainer,
		transactionEventsSubscriber: NewTransactionEventsSubscriber(transactionService),
	}
}

func (container *TransactionEventsContainer) RegisterEvents() {
	container.eventContainer.RegisterEvent(
		CALCULATE_VENDOR_INCOME_EVENT,
		container.transactionEventsSubscriber.SubscribeCalculateVendorIncome,
	)
}
