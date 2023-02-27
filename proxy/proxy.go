package main

import (
	"fmt"
	"sync"
	"time"
)

type Subject interface {
	Method()
	IsDone() bool
}

type RealSubject struct {
	lock   *sync.Mutex
	isDone bool
}

func NewRealSubject() RealSubject {
	return RealSubject{
		lock:   &sync.Mutex{},
		isDone: false,
	}
}

func (s *RealSubject) Method() {
	s.lock.Lock()
	s.isDone = false
	time.Sleep(time.Second * 3)
	fmt.Println("작업 완료")
	s.isDone = true
	s.lock.Unlock()
}

func (s *RealSubject) IsDone() bool {
	return s.isDone
}

type Proxy struct {
	realSubject RealSubject
}

func NewProxySubject() Proxy {
	return Proxy{
		realSubject: NewRealSubject(),
	}
}

func (s *Proxy) Method() {
	fmt.Println("작업 요청중 입니다.")
	go func() {
		s.realSubject.Method()
	}()
}

func (s *Proxy) IsDone() bool {
	return s.realSubject.IsDone()
}
