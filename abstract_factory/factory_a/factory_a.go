package factory_a

import (
	"design-pattern/abstract_factory/component"
	"fmt"
)

type FactoryA struct {
	// Factory
}

func (s *FactoryA) NewObject() component.ObjectType {
	return &ObjectFromA{}
}

func (s *FactoryA) GetDataType0() component.DataType0 {
	return component.NewDataType0("Factory", "A")
}

func (s *FactoryA) SetDataType1(comp component.DataType1) {
	comp.Describe()
}

type ObjectFromA struct{}

func (s *ObjectFromA) Name() {
	fmt.Println("Object generated from factory A")
}
