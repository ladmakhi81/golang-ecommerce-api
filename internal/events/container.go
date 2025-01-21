package events

type Event struct {
	Name    string
	Payload any
}

type eventHandler func(event Event)

type EventsContainer struct {
	events map[string]chan Event
}

func NewEventsContainer2() *EventsContainer {
	return &EventsContainer{
		events: make(map[string]chan Event),
	}
}

func NewEventsContainer() EventsContainer {
	return EventsContainer{
		events: make(map[string]chan Event),
	}
}

func NewEvent(name string, payload any) Event {
	return Event{
		Name:    name,
		Payload: payload,
	}
}

func (container *EventsContainer) RegisterEvent(eventName string, eventHandler eventHandler) {
	if _, exist := container.events[eventName]; !exist {
		container.events[eventName] = make(chan Event)
	}

	// subscriber
	go func() {
		for data := range container.events[eventName] {
			eventHandler(data)
		}
	}()
}

func (container *EventsContainer) PublishEvent(event Event) {
	if channel, channelExistence := container.events[event.Name]; channelExistence {
		channel <- event
	}
}
