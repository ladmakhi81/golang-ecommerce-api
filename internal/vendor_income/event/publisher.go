package vendorincomeevent

import "github.com/ladmakhi81/golang-ecommerce-api/internal/events"

type VendorIncomeEventsPublisher struct {
	eventsContainer *events.EventsContainer
}

func NewVendorIncomeEventsPublisher(
	eventsContainer *events.EventsContainer,
) VendorIncomeEventsPublisher {
	return VendorIncomeEventsPublisher{
		eventsContainer: eventsContainer,
	}
}

func (publisher VendorIncomeEventsPublisher) PublishCalculateVendorIncomeEvent(eventBody CalculateVendorIncomeEventBody) {
	publisher.eventsContainer.PublishEvent(
		events.NewEvent(
			CALCULATE_VENDOR_INCOME_EVENT,
			eventBody,
		),
	)
}
