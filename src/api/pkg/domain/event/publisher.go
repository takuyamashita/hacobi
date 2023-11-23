package event

type EventPublisherIntf[T Event] interface {
	Subscribe(subscribers EventSubscriber[T])
	Publish(event T)
}

type eventPublisherImpl[T Event] struct {
	subscribers []EventSubscriber[T]
}

func NewEventPublisher[T Event]() EventPublisherIntf[T] {
	return &eventPublisherImpl[T]{}
}

func (p *eventPublisherImpl[T]) Subscribe(subscriber EventSubscriber[T]) {

	p.subscribers = append(p.subscribers, subscriber)
}

func (p eventPublisherImpl[T]) Publish(event T) {

	for _, subscriber := range p.subscribers {
		subscriber.Handle(event)
	}
}
