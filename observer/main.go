package main

/*
옵저버 패턴 : 상태 또는 토픽을 발행하는 주제와, 주제를 구독하는 관찰자로 구성하는 디자인 패턴.

옵저버 패턴은 주제(발행)와 관찰자(구독)를 분리하여 느슨한 의존관계를 만든다. 주제는 상태를 발행하지만, 구독하는 관찰자의 내부를 알 필요가 없으며,
주제와 관찰자는 서로에대해 Subject(=Publisher) 와 Observer(=Subscriber) 인터페이스만 알면 된다. 심지어 구독과 해지 여부를 관찰자가 직접 관리하지
않는다면, 관찰자는 주제를 알 필요 자체가 없다. 이를 주제와 관찰자 간에 1:N 의 느슨한 상관관계를 형성한다.

옵저버 패턴과 발행-구독 패턴의 기술적 차이는 동기성 / 비동기성에 있다. 옵저버 패턴은 주제에서 모든 옵저버의 Update 함수를 직접 호출하지만, 발행-구독 패턴은
발행자와 구독자 사이에 브로커 또는 이벤트 버스가 있어 토픽을 푸시하거나, 구독자가 이벤트 버스를 풀링한다. 하지만 이것은 기술적인 차이이고 개념적으로는 동일하다고
볼 수 있다.
*/
func main() {
	subject := ConcretePublisher{}
	subject.Subscribe(&ConcreteSubscriber{Name: "Observer 0"})
	subject.Subscribe(&ConcreteSubscriber{Name: "Observer 1"})
	subject.Subscribe(&ConcreteSubscriber{Name: "Observer 2"})
	subject.Notify()
}
