package main

import "fmt"

func main() {
	var factory Factory
	var product Product
	fmt.Println("--- Factory A ---")
	factory = NewFactoryA()
	product = factory.NewProduct("1")
	product.Describe()

	fmt.Println()
	fmt.Println("--- Factory B ---")
	factory = NewFactoryB()
	product = factory.NewProduct("Y")
	product.Describe()
}
