package contract

type (
	Listener interface {
		Listen(event interface{})
	}

	Event interface {
		Handle()
	}
)
