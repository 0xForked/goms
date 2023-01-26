package notify

import (
	"context"
	"github.com/aasumitro/goms/internal/notify/event"
	"github.com/aasumitro/goms/internal/notify/listener"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func NewNotifyService(redis *redis.Client) {
	ctx := context.Background()

	dispatcher := event.NewDispatcher()
	if err := dispatcher.Register(listener.Listener{}, listener.Notify); err != nil {
		log.Fatalln(err)
	}

	subscriber := redis.Subscribe(ctx, "notify")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		data := []byte(msg.Payload)
		if err := dispatcher.Dispatch(listener.Notify, listener.NotifyEvent{
			Time:    time.Now().UTC(),
			Message: string(data),
		}); err != nil {
			log.Println(err)
		}
	}
}
