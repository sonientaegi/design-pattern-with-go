package main

import "fmt"

/*
중재자를 통한 느슨한 연결

여러 개체를 유기적으로 연결해야 하는 경우, 객체가 객체를 직접 참조하면 의존성이 생기고 유연한 구성이 어려워진다. 수정사항이 생기거나
객체를 추가 또는 삭제하는 경우 전체 코드에 영향을 줄 수 도있다. 이때 객체간의 통신(또는 인터렉션)을 전담하는 중재자를 두고 객체는
오직 중재자 하고만 상호작용하도록 하여 코드의 의존성을 낮추는것이 중재자 패턴이다. 중재자를 통해 상호작용하는 객체를 colleague 로 부른다.

colleague 는 다른 colleague 의 동작을 알 필요가 없다. colleague 는 이벤트나 변경사항은 mediator 로 전달하고, mediator 는
이를 이용해 적절한 작업을 수행하거나, 다른 colleague 에 전달한다. 이벤트를 전달받은 colleague 는 전달 받은 이벤트에 대한 작업을 수행한다.
또 colleague 는 다른 colleague 의 mediator 역할을 할 수 있다.

이는 계층 구조를 갖춘 UI 프로그램에 많이 쓰인다.
*/

func main() {
	mediator := ConcreteMediator{}
	colleagueA := NewConcreteColleagueA(&mediator)
	colleagueB := NewConcreteColleagueB(&mediator)
	mediator.ColleagueA = &colleagueA
	mediator.ColleagueB = &colleagueB

	colleagueA.Callout()
	fmt.Println()
	colleagueB.Callout()
	fmt.Println()
	mediator.Callout()
	fmt.Println()
}
