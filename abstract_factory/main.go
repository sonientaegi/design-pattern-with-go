package main

import (
	"design-pattern/abstract_factory/component"
	"design-pattern/abstract_factory/factory_a"
	"design-pattern/abstract_factory/factory_b"
)

func main() {
	var factory Factory
	var comp0 component.DataType0

	factory = &factory_a.FactoryA{}
	factory.NewObject().Name()
	comp0 = factory.GetDataType0()
	comp0.Describe()
	factory.SetDataType1(component.DataType1{
		Value1: 100,
		Value2: 101,
	})

	factory = &factory_b.FactoryB{}
	factory.NewObject().Name()
	comp0 = factory.GetDataType0()
	comp0.Describe()
	factory.SetDataType1(component.DataType1{
		Value1: 200,
		Value2: 201,
	})

}
