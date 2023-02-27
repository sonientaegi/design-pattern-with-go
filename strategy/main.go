package main

import "fmt"

/*
전략 패턴 : 가장 기본적인 디자인패턴이다.

상속의 목적은 코드 재사용성이 아닌 확장에 있다. 만약 상속을 코드 재사용 목적으로 쓴다면 부모 클래스에 변경때 마다 자식 클래스의 영향 검토를 해야할 것이다.
클래스가 늘어나면 늘어날 수록 변경 검토 양이 늘어나게된다. 또 기존 기능과 다른 동작을 하는 자식 클래스가 필요하다면 상속 & 오버라이드를 해야하고 이것이
계속해서 연쇄되면 클래스 계층 구조가 혼란스럽게 된다. 전략 패턴은 "상속" 대신 "합성"을 이용하여 클래스간의 의존성을 낮추고, 필요한 기능만 구현하도록 한다.

GO 는 상속 대신 임베딩을 하며 임베딩을 이용해서 전략 패턴을 구현할 수 있다.
*/
func main() {
	var target *AbstractObject

	obj0 := NewConcreteObject0()
	obj0.Action()
	target = &obj0.AbstractObject
	target.Action()
	fmt.Println("---")

	obj1 := NewConcreteObject1()
	obj1.Action()
	target = &obj1.AbstractObject
	target.Action()
	fmt.Println("---")

	obj2 := NewConcreteObject2()
	obj2.Action()
	target = &obj2.AbstractObject
	target.Action()
}
