package transactionevent

import "github.com/ladmakhi81/golang-ecommerce-api/internal/events"

type TransactionEventsPublisher struct {
	eventsContainer *events.EventsContainer
}

func NewTransactionEventsPublisher(
	eventsContainer *events.EventsContainer,
) TransactionEventsPublisher {
	return TransactionEventsPublisher{
		eventsContainer: eventsContainer,
	}
}

func (publisher TransactionEventsPublisher) PublishCalculateVendorIncomeEvent(eventBody CalculateVendorIncomeEventBody) {
	publisher.eventsContainer.PublishEvent(
		events.NewEvent(
			CALCULATE_VENDOR_INCOME_EVENT,
			eventBody,
		),
	)
}
