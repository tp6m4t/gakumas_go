package event

import "sync"

type Eventdata interface {
}

type EventBus struct {
	lock        sync.RWMutex
	subscribers map[string][]func(event Eventdata)
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(event Eventdata)),
	}
}

func (eb *EventBus) Publish(eventType string, data Eventdata) {
	eb.Clear(eventType)
	eb.lock.Lock()
	defer eb.lock.Unlock()
	for _, handler := range eb.subscribers[eventType] {
		handler(data)
	}
}

func (eb *EventBus) Clear(eventType string) {
	eb.lock.Lock()
	defer eb.lock.Unlock()
	handlers := eb.subscribers[eventType]
	if len(handlers) == 0 {
		return
	}

	newHandlers := make([]func(event Eventdata), 0, len(handlers))
	for _, handler := range handlers {
		if handler != nil {
			newHandlers = append(newHandlers, handler)
		}
	}
	eb.subscribers[eventType] = newHandlers
}

func (eb *EventBus) Subscribe(eventType string, handler func(event Eventdata)) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}
