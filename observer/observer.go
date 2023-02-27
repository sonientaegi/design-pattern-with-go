package main

type Publisher interface {
	Subscribe(Subscriber)
	Unsubscribe(Subscriber)
	Notify()
}

type Subscriber interface {
	Update(any)
}
