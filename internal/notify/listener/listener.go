package listener

import (
	"log"
)

type Listener struct{}

func (u Listener) Listen(event interface{}) {
	switch event := event.(type) {
	case PushNotifyEvent:
		event.Handle()
	default:
		log.Printf("registered an invalid user event: %T\n", event)
	}
}
