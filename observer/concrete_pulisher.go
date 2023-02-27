package main

type Context struct {
	// blahh
}

type ConcretePublisher struct {
	// Publisher

	subscribers    []Subscriber
	currentContext Context
}

func (s *ConcretePublisher) Subscribe(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

func (s *ConcretePublisher) Unsubscribe(subscriber Subscriber) {
	subscribers := make([]Subscriber, 0)
	for _, sub := range s.subscribers {
		if sub == subscriber {
			continue
		} else {
			subscribers = append(subscribers, sub)
		}
	}
	s.subscribers = subscribers
}

func (s *ConcretePublisher) Notify() {
	for _, subscriber := range s.subscribers {
		subscriber.Update(s.currentContext)
	}
}
