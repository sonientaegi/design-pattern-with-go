package main

import (
	"design-pattern/singleton/singleton"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
싱글턴 : 유일한 객체 생성

싱글턴은 유일한 객체를 생성하는 패턴이다. 목적이나 설계에 따라 프로세스 실행 시 또는 실제 호출시 객체를 생성하도록 생성 시점을 특정할 수도 있고,
멀티 쓰레드 여부에 따라 락메커니즘을 적용할 수 도 있다.

GO 는 단 한번의 로직 호출을 위한 sync.Once 를 제공한다. sync.Once 는 락 메커니즘을 포함한 "유일한 실행"을 보장한다. Do(...) 를 최초 실행한
스레드만 내부 루틴을 수행하며, 그 외 호출한 쓰레드들은 최초의 스레드가 내부 루틴을 종료할 때까지 대기한다. 내부 루틴을 종료한 직후 대기중이던 스레드는
내부 루틴을 수행하지 않고 다음명령을 실행한다. 이로써 sync.Once 는 멀티쓰레드 환경하에서의 싱글턴 패턴의 정합성을 보장할 수 있다.
*/
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	reader := func(wg *sync.WaitGroup) {
		instance := singleton.GetInstance()
		value := instance.GetValue()
		for value < 10 {
			fmt.Println("read", value)
			for {
				newValue := instance.GetValue()
				if newValue == value {
					time.Sleep(time.Millisecond)
				} else {
					value = newValue
					break
				}
			}
		}

		wg.Done()
	}

	writer := func(wg *sync.WaitGroup) {
		instance := singleton.GetInstance()
		for i := 0; i < 10; i++ {

			time.Sleep(time.Duration(float64(time.Second) * rand.Float64()))
			fmt.Println("write")
			instance.Trigger()
		}

		wg.Done()
	}

	go reader(wg)
	go writer(wg)

	wg.Wait()
}
