package events

import (
	"errors"
	"sync"
)

type IEventDispatcher interface {
	Register(eventName string, handler IEventHandler) error
	Has(eventName string, handler IEventHandler) bool
	Dispatch(event IEvent) error
	Unregister(eventName string, handler IEventHandler) error
	UnregisterAll()
}

type EventDispatcher struct {
	handlers map[string][]IEventHandler
}

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

func NewMovementDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]IEventHandler),
	}
}

func (eventDispatcher *EventDispatcher) Register(eventName string, eventHandler IEventHandler) error {
	// check if the event is already registered
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for _, currentHandler := range eventDispatcher.handlers[eventName] {
			if currentHandler == eventHandler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], eventHandler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler IEventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ev *EventDispatcher) Dispatch(event IEvent) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Unregister(eventName string, handler IEventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (eventDispatcher *EventDispatcher) UnregisterAll() {
	eventDispatcher.handlers = make(map[string][]IEventHandler)
}
