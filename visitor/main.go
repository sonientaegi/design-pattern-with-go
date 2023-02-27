package main

/*
객체의 형상을 바꾸지 않고 외부에 기능을 구현하고자 할때 사용한다.

이 경우 캡슐화가 깨지면서 단일책임 원칙에 위배되지만, 변경 사항을 적용하기 위한 공수와 리스크가 이를 능가한다면 비지터 패턴을 고려해볼만하다.
먼저 비지터가 억셉터(대상 객체)를 순회하면서 자신을 인자로 하는 visit 함수를 호출한다. 억셉터는 자신을 인자로 가지는, 자신이 호출해야할 비지터의 accept 함수를
호출한다. 실제 기능 구현은 accept 함수 안에서 이루어진다.
*/
func main() {
	elements := []Acceptable{
		&ElementA{Name: "영"},
		&ElementB{Value: 0},
		&ElementA{Name: "일"},
		&ElementA{Name: "이"},
		&ElementB{Value: 1},
		&ElementA{Name: "삼"},
		&ElementB{Value: 2},
		&ElementB{Value: 3},
	}

	visitor := Visitor{}
	for _, e := range elements {
		e.Accept(visitor)
	}
}
