package main

import (
	"design-pattern/abstract_factory/component"
	"design-pattern/abstract_factory/factory_a"
	"design-pattern/abstract_factory/factory_b"
)

/*
추상 팩토리 패턴 : 객체 전체를 추상화 하여 동일한 인터페이스를 제공한다.

동일한 기능을 제공하는 여러가지 서비스나 제품이 있을때, 코드 변경 없이 설정만으로 사용할 수 있는 호환성을 제공하려면 클라이언트가 구상체에 대한 정보 대신
공통으로 제공하는 명세 즉 인터페이스를 통해 기능을 구현할 수 있어야 한다. 객체를 통채로 추상화 하므로 추상 팩토리 패턴이라 부르며, 이 점에서 팩토리 메소드만
추상화 하는 펙토리 메소드 패턴과 다르다. 이 패턴은 주로 DB 설정처럼 엔진은 서로 달라도 동일한 기능을 제공하는 경우, 그리고 이를 로직 수정 없이
변경할 수 있도록 할때 사용한다.
*/
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
