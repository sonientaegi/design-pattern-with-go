package component

import "fmt"

/* Object type */

type ObjectType interface {
	Name()
}

/* Data type 0 - Read only */

type DataType0 struct {
	value1 string
	value2 string
}

func NewDataType0(value1 string, value2 string) DataType0 {
	return DataType0{
		value1: value1,
		value2: value2,
	}
}

func (s *DataType0) GetValue1() string {
	return s.value1
}

func (s *DataType0) GetValue2() string {
	return s.value2
}

func (s *DataType0) Describe() {
	fmt.Println("Data type 0", s.value1, s.value2)
}

/* Data type 1 - Read & Writable */

type DataType1 struct {
	Value1 int
	Value2 int
}

func (s *DataType1) Describe() {
	fmt.Println("Data type 1", s.Value1, s.Value2)
}
