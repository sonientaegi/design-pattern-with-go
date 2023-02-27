package main

type ConcreteObject0 struct {
	AbstractObject
}

func NewConcreteObject0() ConcreteObject0 {
	return ConcreteObject0{AbstractObject{Capsule: nil}}
}

type ConcreteObject1 struct {
	AbstractObject
}

func NewConcreteObject1() ConcreteObject1 {
	return ConcreteObject1{AbstractObject{Capsule: &ConcreteCapsuleA{}}}
}

type ConcreteObject2 struct {
	AbstractObject
}

func NewConcreteObject2() ConcreteObject2 {
	return ConcreteObject2{AbstractObject{Capsule: &ConcreteCapsuleB{}}}
}
