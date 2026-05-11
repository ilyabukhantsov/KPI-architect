package events

type Event struct {
	Name    string
	Payload interface{}
}

type EventBus struct {
	subscribers []chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{subscribers: make([]chan Event, 0)}
}

func (eb *EventBus) Subscribe() chan Event {
	ch := make(chan Event, 10)
	eb.subscribers = append(eb.subscribers, ch)
	return ch
}

func (eb *EventBus) Publish(name string, payload interface{}) {
	event := Event{Name: name, Payload: payload}
	for _, ch := range eb.subscribers {
		ch <- event
	}
}
