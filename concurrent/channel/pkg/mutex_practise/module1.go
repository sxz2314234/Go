package mutexpractise

import (
	"sync"
)

var (
	X  int64
	Wg sync.WaitGroup
	M  sync.Mutex
)

func Add() {
	for i := 0; i < 5000; i++ {
		M.Lock()
		X++
		M.Unlock()
	}
	Wg.Done()
}
