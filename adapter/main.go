package main

/*
Adapter 는 Target 을 호출하는 Client 의 코드를 변경 하지 않고 Adaptee 와 호환 인터페이스를 제공하는 객체이다.

보통 OOP 언어에서는 이종 객체간의 호환성을 위해 객체어댑터 ( JAVA 처럼 다중상속이 불가한 경우. Adapter 가 Target interface 를 제공) 나
클래스어댑터 ( C++ 처럼 다중 상속이 가능한 경우. Adapter 가 Target 과 Adaptee 를 상속받아 단일 클래스를 생성 ) 를 구현하는데
GO 는 duck type 언어이므로 embedding 을 사용하지 않고도 Adapter 패턴을 구현할 수 있다.

1. Client 가 호출하는 모든 Target 함수를 정리한다.
2. Adaptee 를 필드로 가지는 Adapter 구조체를 정의한다.
3. Adapter 에서 1번의 함수를 모두 mocking 한다.
4. Client 에서 Target 대신 Adapter 로 변수를 교체한다.
*/
func main() {
	target := Target{}
	// target := Adapter{}

	target.Call()
}
