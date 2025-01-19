package vendorincomeevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
)

type VendorIncomeEVentsContainer struct {
	eventContainer               *events.EventsContainer
	vendorIncomeEventsSubscriber VendorIncomeEventsSubscriber
}

func NewVendorIncomeEVentsContainer(
	eventContainer *events.EventsContainer,
	vendorIncomeEventsSubscriber VendorIncomeEventsSubscriber,
) VendorIncomeEVentsContainer {
	return VendorIncomeEVentsContainer{
		eventContainer:               eventContainer,
		vendorIncomeEventsSubscriber: vendorIncomeEventsSubscriber,
	}
}

func (container *VendorIncomeEVentsContainer) RegisterEvents() {
	container.eventContainer.RegisterEvent(
		CALCULATE_VENDOR_INCOME_EVENT,
		container.vendorIncomeEventsSubscriber.SubscribeCalculateVendorIncome,
	)
}
