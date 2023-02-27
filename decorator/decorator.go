package main

import "fmt"

type Decorator struct {
	// Component
	wrappedComponent Component
}

func newDecorator(wrappedComponent Component) Decorator {
	return Decorator{
		wrappedComponent: wrappedComponent,
	}
}

func (s *Decorator) Operator() {
	s.wrappedComponent.Operate()
}

type ConcreteDecoratorA struct {
	Decorator
}

func (s *ConcreteDecoratorA) Operate() {
	fmt.Print("Decorate A IN > ")
	s.wrappedComponent.Operate()
	fmt.Print("Decorate A OUT > ")
}

func NewConcreteDecoratorA(decorated Component) Component {
	return &ConcreteDecoratorA{
		Decorator: newDecorator(decorated),
	}
}

type ConcreteDecoratorB struct {
	Decorator
}

func (s *ConcreteDecoratorB) Operate() {
	fmt.Print("Decorate B IN > ")
	s.wrappedComponent.Operate()
	fmt.Print("Decorate B OUT > ")
}

func NewConcreteDecoratorB(decorated Component) Component {
	return &ConcreteDecoratorB{
		Decorator: newDecorator(decorated),
	}
}
