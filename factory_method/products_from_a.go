package main

import "fmt"

/*
 * Product0 from Factory A
 */

type Product0FromA struct{}

func (s *Product0FromA) Step1() {
	fmt.Println("Step 1")
}

func (s *Product0FromA) Step2() {
	fmt.Println("Step 2")
}

func (s *Product0FromA) Describe() {
	fmt.Println("Product 0 from Factory A")
}

/*
 * Product1 from Factory A
 */

type Product1FromA struct{}

func (s *Product1FromA) Step1() {
	fmt.Println("Step 1")
}

func (s *Product1FromA) Step2() {
	fmt.Println("Step 2")
}

func (s *Product1FromA) Describe() {
	fmt.Println("Product 1 from Factory A")
}
