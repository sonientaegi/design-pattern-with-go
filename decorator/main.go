package main

import "fmt"

/*
데코레이터 패턴 : 또는 장식자 패턴. 기존 객체 변경 없이, 데코레이터(또는 장식자)로 감싸 새로운 기능을 추가한다.

특정한 동작을 하는 객체가 존재할 때, 이 객체를 변경하지 않고 기존 객체와 동일한 인터페이스를 구현한 장식자로 객체를 감싸 기능을 추가하는 방법이다.
클라이언트 입장에서는 인터페이스가 동일하므로 코드 변경을 알 필요가 없으며, 기존 객체를 수정할 필요가 없으므로 구현의 복잡도도 감소한다.
다만 장식자를 장식자로, 다시 또 다른 장식자로 감싸다보면 계층구조가 비대해지는 단점이 있다.
*/
func main() {
	wrappedComponent := NewConcreteDecoratorB(NewConcreteDecoratorA(&ConcreteComponent{Description: "Hello!"}))
	wrappedComponent.Operate()
	fmt.Println("DONE")
}
