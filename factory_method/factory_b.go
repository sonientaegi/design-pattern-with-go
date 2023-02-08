package main

import "strings"

type FactoryB struct {
	FactoryAbstract
}

func NewFactoryB() Factory {
	factory := FactoryB{}
	factory.FactoryAbstract.newProduct = factory.newProduct

	return &factory
}

func (s *FactoryB) newProduct(kind string) Product {
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
