package main

import "fmt"

type Colleague interface {
	OnEvent(Event, string)
}

type AbstractColleague struct {
	Colleague
	mediator Mediator
}

type ConcreteColleagueA struct {
	AbstractColleague
}

func NewConcreteColleagueA(mediator Mediator) ConcreteColleagueA {
	return ConcreteColleagueA{
		AbstractColleague: AbstractColleague{
			mediator: mediator,
		},
	}
}

func (s *ConcreteColleagueA) OnEvent(event Event, context string) {
	fmt.Println("Colleague A", event, context)
}

func (s *ConcreteColleagueA) Callout() {
	fmt.Println("Colleague A is called!")
	s.mediator.SetEvent(CalloutA, "from extern")
}

type ConcreteColleagueB struct {
	AbstractColleague
}

func NewConcreteColleagueB(mediator Mediator) ConcreteColleagueB {
	return ConcreteColleagueB{
		AbstractColleague: AbstractColleague{
			mediator: mediator,
		},
	}
}

func (s *ConcreteColleagueB) OnEvent(event Event, context string) {
	fmt.Println("Colleague B", event, context)
}

func (s *ConcreteColleagueB) Callout() {
	fmt.Println("Colleague B is called!")
	s.mediator.SetEvent(CalloutB, "from extern")
}
