package listener

import (
	"github.com/aasumitro/goms/internal/notify/event"
	"log"
	"time"
)

const Notify event.Name = "notify.push"

type NotifyEvent struct {
	Time    time.Time
	Message string
}

func (e NotifyEvent) Handle() {
	log.Printf("creating: %+v\n", e)
}
