package vendorincomeevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
)

type VendorIncomeEventsContainer struct {
	eventContainer               *events.EventsContainer
	vendorIncomeEventsSubscriber VendorIncomeEventsSubscriber
}

func NewVendorIncomeEventsContainer(
	eventContainer *events.EventsContainer,
	vendorIncomeEventsSubscriber VendorIncomeEventsSubscriber,
) VendorIncomeEventsContainer {
	return VendorIncomeEventsContainer{
		eventContainer:               eventContainer,
		vendorIncomeEventsSubscriber: vendorIncomeEventsSubscriber,
	}
}

func (container *VendorIncomeEventsContainer) RegisterEvents() {
	container.eventContainer.RegisterEvent(
		events.CALCULATE_VENDOR_INCOME_EVENT,
		container.vendorIncomeEventsSubscriber.SubscribeCalculateVendorIncome,
	)
}
