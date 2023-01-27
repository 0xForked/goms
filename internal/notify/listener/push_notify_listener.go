package listener

import (
	"github.com/aasumitro/goms/internal/notify/event"
	"log"
	"time"
)

const PushNotify event.Name = "notify.push"

type PushNotifyEvent struct {
	Time    time.Time
	Message string
}

func (e PushNotifyEvent) Handle() {
	log.Printf("creating: %+v\n", e)
}
