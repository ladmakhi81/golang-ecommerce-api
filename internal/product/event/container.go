package productevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
)

type ProductEventsContainer struct {
	eventContainer          *events.EventsContainer
	productEventsSubscriber ProductEventsSubscriber
}

func NewProductEventsContainer(
	eventContainer *events.EventsContainer,
	productEventsSubscriber ProductEventsSubscriber,
) ProductEventsContainer {
	return ProductEventsContainer{
		eventContainer:          eventContainer,
		productEventsSubscriber: productEventsSubscriber,
	}
}

func (container *ProductEventsContainer) RegisterEvents() {
	container.eventContainer.RegisterEvent(
		events.PRODUCT_CREATED_EVENT,
		container.productEventsSubscriber.SubscribeProductCreated,
	)

	container.eventContainer.RegisterEvent(
		events.PRODUCT_VERIFIED_EVENT,
		container.productEventsSubscriber.SubscribeProductVerification,
	)
}
