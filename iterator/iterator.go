package main

import "math/rand"

type Iterable interface {
	Iterator() chan any
}

type RandomIntGenerator struct {
	Count int
}

func (s RandomIntGenerator) Iterator() chan any {
	iter := make(chan any, s.Count)
	for i := 0; i < s.Count; i++ {
		iter <- rand.Int()
	}
	close(iter)

	return iter
}
