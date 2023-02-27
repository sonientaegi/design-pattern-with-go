package main

import "fmt"

type State interface {
	OperateA()
	OperateB()
	OperateC()
}

/* === State A === */

type StateA struct {
	stateMachine *StateMachine
}

func (s *StateA) OperateA() {
	fmt.Println("A -> B")
	s.stateMachine.SetCurrentState(StateTypeB)
}

func (s *StateA) OperateB() {
	fmt.Println("A -> C")
	s.stateMachine.SetCurrentState(StateTypeC)
}

func (s *StateA) OperateC() {
	fmt.Println("A -> D")
	s.stateMachine.SetCurrentState(StateTypeD)
}

/* === State B === */

type StateB struct {
	stateMachine *StateMachine
}

func (s *StateB) OperateA() {
	fmt.Println("B -> C")
	s.stateMachine.SetCurrentState(StateTypeC)
}

func (s *StateB) OperateB() {
	fmt.Println("B -> B")
	s.stateMachine.SetCurrentState(StateTypeB)
}

func (s *StateB) OperateC() {
	fmt.Println("B -> D")
	s.stateMachine.SetCurrentState(StateTypeD)
}

/* === State C === */

type StateC struct {
	stateMachine *StateMachine
}

func (s *StateC) OperateA() {
	fmt.Println("C -> D")
	s.stateMachine.SetCurrentState(StateTypeD)
}

func (s *StateC) OperateB() {
	fmt.Println("C -> B")
	s.stateMachine.SetCurrentState(StateTypeB)
}

func (s *StateC) OperateC() {
	panic("Invalid operation")
}

/* === State D === */

type StateD struct {
	stateMachine *StateMachine
}

func (s *StateD) OperateA() {
	fmt.Println("D -> A")
	s.stateMachine.SetCurrentState(StateTypeA)
}

func (s *StateD) OperateB() {
	fmt.Println("D -> B")
	s.stateMachine.SetCurrentState(StateTypeB)
}

func (s *StateD) OperateC() {
	fmt.Println("D -> C")
	s.stateMachine.SetCurrentState(StateTypeC)
}
