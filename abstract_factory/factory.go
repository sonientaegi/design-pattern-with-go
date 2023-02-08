package main

import "design-pattern/abstract_factory/component"

type Factory interface {
	NewObject() component.ObjectType
	GetDataType0() component.DataType0
	SetDataType1(comp component.DataType1)
}
