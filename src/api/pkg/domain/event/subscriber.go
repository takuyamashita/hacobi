package event

type EventSubscriber[T Event] interface {
	Handle(event T)
}
