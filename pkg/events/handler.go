package events

import "sync"

type IEventHandler interface {
	Handle(event IEvent, wg *sync.WaitGroup)
}
