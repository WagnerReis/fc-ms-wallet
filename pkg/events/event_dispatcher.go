package events

import (
	"errors"
	"slices"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already been registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if handlers, ok := ed.handlers[eventName]; ok {
		if slices.Contains(handlers, handler) {
			return true
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := ed.handlers[eventName]; ok {
		index := slices.Index(handlers, handler)
		ed.handlers[eventName] = append(handlers[:index], handlers[index+1:]...)
	}
	return nil
}
