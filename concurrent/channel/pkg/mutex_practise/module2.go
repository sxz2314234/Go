package mutexpractise

import (
	"fmt"
	"sync"
	"time"
)

var (
	Y   = 0
	Wg2 sync.WaitGroup
	// lock sync.Mutex
	Rwlock sync.RWMutex
)

func Read() {
	defer Wg2.Done()

	// lock.Lock()
	Rwlock.RLock()
	fmt.Println(Y)
	time.Sleep(time.Millisecond)
	Rwlock.RUnlock()
	// lock.Unlock()
}

func Write() {
	defer Wg2.Done()
	Rwlock.Lock()
	// lock.Lock()
	Y += 1
	time.Sleep(time.Millisecond * 5)
	Rwlock.Unlock()
	// lock.Unlock()
}
