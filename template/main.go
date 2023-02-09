package main

import "fmt"

/*
공통 로직과 개별 로직을 분리한다.

템플릿 패턴은 공통 로직과 개별 로직을 분리한 뒤 각 구현체에서 개별 로직 부분만 (재)정의하여 사용한다. 모든 구현체에 대하여 공통의
로직이나 구성이 존재하고, 각 구현체 별로 (재)정의가 필요한 부분을 분리할 수 있을때 템플릿 패턴을 쓸 수 있다. 마치 괄호가 있는
문장안에 필요한 단어나 숫자를 넣어 문장을 완성하는것과 같은 모습이다. 템플릿 구조체는 공통 함수와 (재)정의 함수를 이용하여
공통 로직을 구현한다. 그리고 템플릿 구조체를 임베딩하는 구조체에서 필요한 부분을 (재)정의 한다.

OOP 언어에서 탬플릿 패턴을 구현할때, 반드시 정의가 필요한 함수는 추상함수로 선언한다. 반면 필요시에만 구현이 필요한 함수의 경우
탬플릿 클래스 안에 원형을 선언해두고 확장 클래스에서 재정의해 사용하는데 이를 후킹 함수라고 한다.

Go 는 Virtual 함수를 지원하지 않으므로 재정의를 통한 후킹 구현은 불가능하다. 대신...
1. 후킹용 함수 포인터를 템플릿 구조체에 선언하고
2. 템플릿 생성자에서 기본 함수를 후킹용 함수 포인터에 연결한뒤
3. 구현체에서 필요시 이를 재정의 하여 사용
하는 방법으로 구현할 수 있다. 만약 후킹을 사용하지 않는다면 예제 코드의 템플릿 생성자 NewTemplate() 은 불필요 하다.
*/

func main() {
	fmt.Println("=== Template extension A ===")
	extA := NewExtendA()
	extA.Execute()
	fmt.Println()

	fmt.Println("=== Template extension B ===")
	extB := NewExtendB()
	extB.Execute()
	fmt.Println()

	fmt.Println("=== Template extension C with hook overriding ===")
	extC := NewExtendC()
	extC.Execute()
	fmt.Println()

	fmt.Println("=== Panic by calling abstract method ===")
	tmpl := NewTemplate()
	tmpl.Execute()
	fmt.Println()
}
