package main

import (
	"fmt"
	"time"
)

/*
Proxy 는 클라이언트가 실재 객체를 호출하는 대신 대리 객체를 통해 요청을 처리하는 디자인패턴이다.

Proxy 와 RealSubject 는 동일한 Subject 인터페이스를 구현한다. 클라이언트는 RealSubject 대신 Proxy 를 생성하고 이 Proxy 에 요청을 보낸다.
Proxy 는 내부에 RealSubject 를 가지고 있으며 Proxy 를 통한 요청의 실질적인 처리는 RealSubject 가 수행한다. Proxy 는 유형에 따라 크게 세가지로
나뉜다.
- 원격 프록시 : 원격 프로세스를 로컬 오브젝트처럼 호출.
- 가상 프록시 : 기능을 대리 수행.
- 보호 프록시 : 원래 기능에 접근 보호 제공 (Read or Write or ReadWrite)

다음 예제는 Main Thread 블로킹 없이 Heavy Job 을 수행하는 가상 프록시이다.
*/
func main() {
	subject := NewProxySubject()
	subject.Method()
	for !subject.IsDone() {
		fmt.Println("메인 스레드는 열일 중")
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("메인 스레드는 열일 중")
}
