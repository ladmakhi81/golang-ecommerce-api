package orderevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
)

type OrderEventsContainer struct {
	eventContainer        *events.EventsContainer
	orderEventsSubscriber OrderEventsSubscriber
}

func NewOrderEventsContainer(
	eventContainer *events.EventsContainer,
	orderEventsSubscriber OrderEventsSubscriber,
) OrderEventsContainer {
	return OrderEventsContainer{
		eventContainer:        eventContainer,
		orderEventsSubscriber: orderEventsSubscriber,
	}
}

func (container *OrderEventsContainer) RegisterEvents() {
	container.eventContainer.RegisterEvent(
		events.CHANGED_ORDER_STATUS_EVENT,
		container.orderEventsSubscriber.SubscribeChangeOrderStatus,
	)
}
