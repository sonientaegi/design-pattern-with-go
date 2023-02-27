package main

import "fmt"

type Component interface {
	Add(component Component)
	Remove(component Component)
	Operate()
}

type Composite struct {
	Name   string
	leaves []Component
}

func (s *Composite) Add(component Component) {
	// append(nil, component) -> 새로운 slice
	s.leaves = append(s.leaves, component)
}

func (s *Composite) Remove(component Component) {
	s.Remove(component)
}

func (s *Composite) Operate() {
	fmt.Println("Operate : ", s.Name)
	// range slice type nil -> Do nothing
	for _, component := range s.leaves {
		component.Operate()
	}
}

type Leaf struct {
	Name string
}

func (s *Leaf) Add(component Component) {
	// Do nothing
}

func (s *Leaf) Remove(component Component) {
	// Do nothing
}

func (s *Leaf) Operate() {
	fmt.Println("Operate : ", s.Name)
}
