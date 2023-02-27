package main

import (
	"math/rand"
	"time"
)

type Sender struct {
	invoker    *Invoker
	commandSet []Command
}

type ConcreteCommand struct {
	receiver *Receiver
}

func (s *ConcreteCommand) Execute() {
	s.receiver.DoSomething()
}

type InvokerTeminatCommand struct {
	invoker *Invoker
}

func (s *InvokerTeminatCommand) Execute() {
	s.invoker.Stop()
}

func (s *Sender) Run() {
	go func() {
		for s.invoker.SetCommand(s.commandSet[rand.Intn(len(s.commandSet))]) {
			time.Sleep(time.Second)
		}
	}()
}

/*
Command 패턴은 명령을 보내는 주체와 명령을 수행하는 주체를 분리하는 패턴이다.

Command 패턴은 명령을 실제로 구현하는 Receiver 객체와 작업 명세를 정의하는 Command 인터페이스, 그리고 작업을 관리하고 실행하는 Invoker 객체로
구성한다. 실행해야할 명령은 Invoker 에 등록하고 Invoker 는 이를 실행한다. 이때 Invoker 는 Command 인터페이스에 대해서만 알 뿐이며
그 내용 - Receiver 가 누구인지, 어떤 명령을 수행하는지 - 는 알 필요가 없다. 만약 모든 명령을 명시적으로 수행하는 객체를 구현한다면, 명령을 추가할때마다
이 객체에 구현을 추가해야한다. 하지만 Command 패턴을 구현한다면 명령을 요청하는 객체와 명령을 실행하는 객체를 분리할 수 있으므로 명령을 실행하는 Invoker 는
변경이 필요 없게된다.

이 패턴은 주로 멀티스레드 환경에서 많이 사용한다. 안드로이드는 Looper 를 이용해 작업 스레드의 요청을 메인스레드에서 수행하도록 하며, 윈도우 프로그래밍에서는
Message Pump 를 구현하여 각종 이벤트를 방송한다. 멀티 워커를 가지는 웹서버의 경우 명령이 들어오면 유휴 상태의 워커가 커맨드를 가지고와 처리하고 완료후 다시
유휴 상태에 들어가기를 반복한다.
*/
func main() {
	invoker := Invoker{
		commands: make(chan Command, 100),
	}

	receiverA := Receiver{Name: "receiver A"}
	receiverB := Receiver{Name: "receiver B"}
	receiverC := Receiver{Name: "receiver C"}

	commandA := ConcreteCommand{receiver: &receiverA}
	commandB := ConcreteCommand{receiver: &receiverB}
	commandC := ConcreteCommand{receiver: &receiverC}
	commandX := InvokerTeminatCommand{invoker: &invoker}

	sender := Sender{
		invoker: &invoker,
		commandSet: []Command{
			&commandA, &commandB, &commandC,
			&commandA, &commandB, &commandC,
			&commandA, &commandB, &commandC,
			&commandA, &commandB, &commandC,
			&commandA, &commandB, &commandC,
			&commandA, &commandB, &commandC,
			&commandX},
	}
	sender.Run()
	invoker.Run()
}
