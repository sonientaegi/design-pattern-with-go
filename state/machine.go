package main

import "math/rand"

type StateType int

const (
	StateTypeA StateType = iota
	StateTypeB
	StateTypeC
	StateTypeD
)

type StateMachine struct {
	currentState State
	states       [4]State
}

func NewStateMachine() *StateMachine {
	stateMachine := new(StateMachine)
	stateMachine.states = [4]State{
		&StateA{stateMachine: stateMachine},
		&StateB{stateMachine: stateMachine},
		&StateC{stateMachine: stateMachine},
		&StateD{stateMachine: stateMachine},
	}
	stateMachine.SetCurrentState(StateTypeA)
	return stateMachine
}

func (s *StateMachine) Operate() {
	operate := rand.Intn(99)
	switch operate % 3 {
	case 0:
		s.currentState.OperateA()
	case 1:
		s.currentState.OperateB()
	case 2:
		s.currentState.OperateC()
	}
}

func (s *StateMachine) SetCurrentState(stateType StateType) {
	s.currentState = s.states[stateType]
}
