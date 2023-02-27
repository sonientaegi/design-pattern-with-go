package main

/*
컴포지트 패턴

구성요소의 종류와 상관없이 동일한 인터페이스를 제공하여 계층 관계를 만드는 디자인 패턴이다. 컴포지트 패턴을 이용하면 노드와 리프를 이용한 트리구조를 만들때
별도의 구분 없이 동일한 인터페이스를 제공하여 클라이언트가 별도의 분기를 타지 않아도 모든 객체에 접근할 수 있게 만들 수 있다. 주로 반복자 (Iterator) 와
함께 사용하여 순회 (traversal) 를 구현한다.
*/
func main() {
	cat110 := Composite{Name: "CAT 110"}
	cat110.Add(&Leaf{Name: "ITEM 11a"})
	cat110.Add(&Leaf{Name: "ITEM 11b"})
	cat110.Add(&Leaf{Name: "ITEM 11c"})

	cat120 := Composite{Name: "CAT 120"}
	cat120.Add(&Leaf{Name: "ITEM 12a"})
	cat120.Add(&Leaf{Name: "ITEM 12b"})
	cat120.Add(&Leaf{Name: "ITEM 12c"})

	cat10 := Composite{Name: "CAT 10"}
	cat10.Add(&cat110)
	cat10.Add(&cat120)
	cat10.Add(&Leaf{Name: "ITEM 1a"})
	cat10.Add(&Leaf{Name: "ITEM 1b"})

	cat210 := Composite{Name: "CAT 210"}
	cat210.Add(&Leaf{Name: "ITEM 21a"})
	cat210.Add(&Leaf{Name: "ITEM 21b"})
	cat210.Add(&Leaf{Name: "ITEM 21c"})

	cat20 := Composite{Name: "CAT 20"}
	cat20.Add(&cat210)

	cat0 := Composite{Name: "CAT 0"}
	cat0.Add(&Leaf{Name: "ITEM a"})
	cat0.Add(&cat10)
	cat0.Add(&cat20)

	cat0.Operate()

}
