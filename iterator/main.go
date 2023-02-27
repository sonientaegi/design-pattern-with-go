package main

import "fmt"

/*
GO 는 개발자가 직접 정의할 수 있는 Iterator 인터페이스를 제공하지 않는다. built-in 배열, 맵 과 chan 만 range 의 피연산자로 사용 가능하다.
이를 우회하기 위해 사용할 수 있는 방법이 몇가지 있다.
- 배열 사본 반환하기.
- Next() (any, bool) 함수 구현하기.
- go routine 과 channel 로 구현하기.
- channel 반환하기.

여기서는 이 중 channel 반환하기 예시를 들어본다. channel 은 close 하기 전 & channel size 에 도달할 때 까지 wait 없이 send 가 가능하다.
또 이미 close 를 해도 channel 에 남아있는 데이터를 receive 할 수 있다. 이를 이용해 데이터 개수만큼의 크기를 가지는 channel 을 선언하고 데이터를
모두 send 다음 close 한 뒤 이를 반환하면 반복자로 사용할 수 있다.

별도의 go routine 없이 생성할 수 있다는 장점이 있지다. 하지만 데이터를 미리 체널에 저장하는 만큼 메모리를 차지하며, iterator 를 재사용 할 수 없다.
*/
func main() {
	for i := range (RandomIntGenerator{Count: 10}.Iterator()) {
		fmt.Println(i)
	}
}
