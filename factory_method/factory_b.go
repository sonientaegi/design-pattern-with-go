package main

import "strings"

type FactoryB struct {
	AbstractFactory
}

func NewFactoryB() Factory {
	factory := FactoryB{}
	factory.AbstractFactory.abstractedFactoryMethod = factory.factoryMethod

	return &factory
}

func (s *FactoryB) factoryMethod(kind string) Product {
	var product Product
	switch strings.ToLower(kind) {
	case "x":
		product = &ProductXFromB{}
	case "y":
		product = &ProductYFromB{}
	default:
		panic(kind + " is not proper product type.")
	}
	return product
}
