package main

import "fmt"

type ConcreteCapsuleA struct {
	// Encapsulate
}

func (s *ConcreteCapsuleA) Method() {
	fmt.Println("Call Method type A")
}

type ConcreteCapsuleB struct {
	// Encapsulate
}

func (s *ConcreteCapsuleB) Method() {
	fmt.Println("Call Method type B")
}
