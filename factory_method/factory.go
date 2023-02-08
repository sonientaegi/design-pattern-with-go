package main

type Factory interface {
	NewProduct(kind string) Product
}

type FactoryAbstract struct {
	newProduct func(string) Product
}

func (s FactoryAbstract) NewProduct(kind string) Product {
	product := s.newProduct(kind)
	product.Step1()
	product.Step2()
	return product
}
