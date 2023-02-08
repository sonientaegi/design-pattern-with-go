package factory_b

import (
	"design-pattern/abstract_factory/component"
	"fmt"
)

type FactoryB struct {
	// Factory
}

func (s *FactoryB) NewObject() component.ObjectType {
	return &ObjectFromB{}
}

func (f *FactoryB) GetDataType0() component.DataType0 {
	return component.NewDataType0("Factory", "B")
}

func (f *FactoryB) SetDataType1(comp component.DataType1) {
	comp.Describe()
}

type ObjectFromB struct{}

func (s *ObjectFromB) Name() {
	fmt.Println("Object generated from factory B")
}
