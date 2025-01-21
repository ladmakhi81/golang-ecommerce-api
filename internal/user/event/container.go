package userevent

import (
	"github.com/ladmakhi81/golang-ecommerce-api/internal/events"
)

type UserEventsContainer struct {
	eventContainer       *events.EventsContainer
	userEventsSubscriber UserEventsSubscriber
}

func NewUserEventsContainer(
	eventContainer *events.EventsContainer,
	userEventsSubscriber UserEventsSubscriber,
) UserEventsContainer {
	return UserEventsContainer{
		eventContainer:       eventContainer,
		userEventsSubscriber: userEventsSubscriber,
	}
}

func (container *UserEventsContainer) RegisterEvents() {
	container.eventContainer.RegisterEvent(
		events.USER_REGISTERED_EVENT,
		container.userEventsSubscriber.SubscribeUserRegistered,
	)

	container.eventContainer.RegisterEvent(
		events.USER_COMPLETE_PROFILE_EVENT,
		container.userEventsSubscriber.SubscribeCompleteProfile,
	)

	container.eventContainer.RegisterEvent(
		events.USER_VERIFIED_EVENT,
		container.userEventsSubscriber.SubscribeUserVerification,
	)
}
