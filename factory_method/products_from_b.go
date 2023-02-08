package main

import "fmt"

/*
 * ProductX from Factory A
 */

type ProductXFromB struct{}

func (s *ProductXFromB) Step1() {
	fmt.Println("Step 1")
}

func (s *ProductXFromB) Step2() {
	fmt.Println("Step 2")
}

func (s *ProductXFromB) Describe() {
	fmt.Println("Product X from Factory B")
}

/*
 * ProductY from Factory A
 */

type ProductYFromB struct{}

func (s *ProductYFromB) Step1() {
	fmt.Println("Step 1")
}

func (s *ProductYFromB) Step2() {
	fmt.Println("Step 2")
}

func (s *ProductYFromB) Describe() {
	fmt.Println("Product Y from Factory B")
}
