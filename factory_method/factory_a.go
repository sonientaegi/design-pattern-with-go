package main

import "strings"

type FactoryA struct {
	FactoryAbstract
}

func NewFactoryA() Factory {
	factory := FactoryA{}
	factory.FactoryAbstract.newProduct = factory.newProduct

	return &factory
}

func (s *FactoryA) newProduct(kind string) Product {
	var product Product
	switch strings.ToLower(kind) {
	case "0":
		product = &Product0FromA{}
	case "1":
		product = &Product1FromA{}
	default:
		panic(kind + " is not proper product type.")
	}
	return product
}
