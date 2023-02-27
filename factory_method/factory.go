package main

type Factory interface {
	FactoryMethod(string) Product
}

type AbstractFactory struct {
	// Factory
	abstractedFactoryMethod func(string) Product
}

func (s AbstractFactory) FactoryMethod(kind string) Product {
	product := s.abstractedFactoryMethod(kind)
	product.Step1()
	product.Step2()
	return product
}
