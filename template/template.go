package main

import "fmt"

type Template struct {
	AbstractMethod func()
	HookMethod     func()
}

// NewTemplate 은 Template 구조체의 생성자이다. 만약 후킹을 쓰지 않거나, 별도의 초기화가 필요하지 않다면
// Template 구조체의 생성자는 불필요하다.
func NewTemplate() Template {
	instance := Template{}
	instance.HookMethod = instance.OriginalHookMethod

	return instance
}

func (s *Template) OriginalHookMethod() {
	fmt.Println("This is original hook method.")
}

func (s *Template) CommonMethod0() {
	fmt.Println("This is common logic 0.")
}

func (s *Template) CommonMethod1() {
	fmt.Println("This is common logic 1.")
}

func (s *Template) Execute() {
	s.CommonMethod0()
	s.AbstractMethod()
	s.HookMethod()
	s.CommonMethod1()
}
