package main

import "fmt"

type Event int

const (
	CalloutA Event = iota
	CalloutB
	CalloutBoth
)

type Mediator interface {
	SetEvent(Event, string)
}

type ConcreteMediator struct {
	ColleagueA Colleague
	ColleagueB Colleague
}

func (s *ConcreteMediator) SetEvent(event Event, context string) {
	switch event {
	case CalloutA:
		s.ColleagueB.OnEvent(CalloutA, context)
	case CalloutB:
		s.ColleagueA.OnEvent(CalloutB, context)
	}
}

func (s *ConcreteMediator) Callout() {
	fmt.Println("Mediator is called!")
	context := "by Event broadcaster"
	s.ColleagueA.OnEvent(CalloutBoth, context)
	s.ColleagueB.OnEvent(CalloutBoth, context)
}
