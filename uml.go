package design_pattern

type Interface interface {
	Method()
}

type AbstractStruct struct {
	Method func()
}

type Property struct {
	Value any
}

type Struct struct {
	AbstractStruct
	property Property
}

func (s Struct) Method() {
	//TODO implement me
	panic("implement me")
}

func main() {
	_ = Struct{
		AbstractStruct: AbstractStruct{
			Method: func() {
				// Do something
			},
		},
		property: Property{
			Value: nil,
		},
	}
}
