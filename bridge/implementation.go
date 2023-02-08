package main

import "fmt"

type Implementor[T any] interface {
	MethodA(value T)
	MethodB() T
}

type Implement0[T any] struct {
	// Implementor[T]

	Value T
}

func (s Implement0[T]) MethodA(value T) {
	fmt.Println("Call Implement0.MethodA with ", value)
}

func (s Implement0[T]) MethodB() T {
	fmt.Println("Call Implement0.MethodB")
	return s.Value
}

type Implement1[T any] struct {
	// Implementor[T]

	Value T
}

func (s Implement1[T]) MethodA(value T) {
	fmt.Println("Call Implement1.MethodA with ", value)
}

func (s Implement1[T]) MethodB() T {
	fmt.Println("Call Implement1.MethodB")
	return s.Value
}
