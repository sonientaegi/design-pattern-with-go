package main

import "fmt"

/*
기능과 구현을 분리한다.

브릿지 패턴은 abstraction(기능)과 implementation(구현)으로 분리한다. 기능을 abstraction 이라고 하지만 실제 추상클래스를 뜻하는게 아니라
기능의 일부(또는 대부분)을 추상화하여 실제 구현과 분리시키기 때문에 기능부분을 abstraction 이라고한다. 기능은 기능대로 상속 계층 구조를, 구현은 구현대로
상속 계층 구조를 가져갈 수 있다. 단 Go 는 상속이 아닌 embedding 이므로 계층 구조를 가져가려면 interface 또는 Abstract Struct 를 이용해야한다.

브릿지 패턴은 하나의 기능을 여러가지의 알고리즘으로 구현하고자 할 때 유용하다.
*/
func main() {
	fmt.Println("--- Abstraction struct with implementor 0 ---")
	funcWithImpl0 := &Abstraction[int]{
		Impl: Implement0[int]{
			Value: 2,
		},
	}
	funcWithImpl0.CallMethodA(1)
	funcWithImpl0.CallMethodB()

	fmt.Println()
	fmt.Println("--- Abstraction struct with implementor 1 ---")
	funcWithImpl1 := &Abstraction[int]{
		Impl: Implement1[int]{
			Value: 20,
		},
	}
	funcWithImpl1.CallMethodA(10)
	funcWithImpl1.CallMethodB()

	fmt.Println()
	fmt.Println("--- Abstraction Sub struct with implementor 1 ---")
	funcSubstruct := &AbstractionSub0[float64]{}
	funcSubstruct.Impl = Implement1[float64]{
		Value: 3.14,
	}
	funcSubstruct.CallBothMethod(6.09)
}
