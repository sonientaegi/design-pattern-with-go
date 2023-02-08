package main

import (
	"design-pattern/singleton/singleton"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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
