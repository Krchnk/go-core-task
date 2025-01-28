package customwaitgroup

import (
	"sync"
)

type CustomWaitGroup struct {
	counter int
	mu      sync.Mutex
	done    chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		done: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter += delta
	if wg.counter < 0 {
		panic("negative counter")
	}
	if wg.counter == 0 {
		close(wg.done)
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	if wg.counter == 0 {
		wg.mu.Unlock()
		return
	}
	wg.mu.Unlock()

	<-wg.done
}
