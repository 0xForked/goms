package event

import (
	"fmt"
	"github.com/aasumitro/goms/internal/notify/domain/contract"
)

type Name string

type job struct {
	eventName Name
	eventType interface{}
}

type Dispatcher struct {
	jobs   chan job
	events map[Name]contract.Listener
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		jobs:   make(chan job),
		events: make(map[Name]contract.Listener),
	}

	go d.consume()

	return d
}

func (d *Dispatcher) Register(listener contract.Listener, names ...Name) error {
	for _, name := range names {
		if _, ok := d.events[name]; ok {
			return fmt.Errorf("the '%s' event is already registered", name)
		}

		d.events[name] = listener
	}

	return nil
}

func (d *Dispatcher) Dispatch(name Name, event interface{}) error {
	if _, ok := d.events[name]; !ok {
		return fmt.Errorf("the '%s' event is not registered", name)
	}

	d.jobs <- job{eventName: name, eventType: event}

	return nil
}

func (d *Dispatcher) consume() {
	for job := range d.jobs {
		d.events[job.eventName].Listen(job.eventType)
	}
}
