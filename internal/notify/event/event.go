package event

type Name string

type Event interface {
	Handle()
}
