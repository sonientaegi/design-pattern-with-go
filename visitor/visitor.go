package main

import "fmt"

type Visitable interface {
	VisitElementA(a *ElementA)
	VisitElementB(a *ElementB)
}

type Acceptable interface {
	Accept(Visitable)
}

type Visitor struct {
	// Visitable
}

func (_ Visitor) VisitElementA(element *ElementA) {
	fmt.Println("Element name is", element.Name)
}

func (_ Visitor) VisitElementB(element *ElementB) {
	fmt.Println("Element value is", element.Value)
}

type ElementA struct {
	Name string
}

func (s *ElementA) Accept(visitable Visitable) {
	visitable.VisitElementA(s)
}

type ElementB struct {
	Value int
}

func (s *ElementB) Accept(visitable Visitable) {
	visitable.VisitElementB(s)
}
