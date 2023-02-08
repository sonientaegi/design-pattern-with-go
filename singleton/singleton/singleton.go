package singleton

import "sync"

var instance *Singleton = nil
var once = &sync.Once{}

func GetInstance() *Singleton {
	if instance == nil {
		// DCL
		once.Do(func() {
			instance = &Singleton{
				mtx:   &sync.Mutex{},
				value: 0,
			}
		})
	}

	return instance
}

type Singleton struct {
	mtx   *sync.Mutex
	value int
}

func (s *Singleton) Trigger() {
	s.value++
}

func (s *Singleton) GetValue() int {
	return s.value
}
