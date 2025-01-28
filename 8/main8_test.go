package customwaitgroup

import (
	"sync"
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
	}()

	wg.Wait()

	wg = NewCustomWaitGroup()
	const numTasks = 5
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}

	wg.Wait()

	wg = NewCustomWaitGroup()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when counter goes negative")
		}
	}()

	wg.Add(1)
	wg.Done()
	wg.Done()
}

func TestCustomWaitGroupConcurrent(t *testing.T) {
	wg := NewCustomWaitGroup()
	const numTasks = 100
	wg.Add(numTasks)

	var counter int
	var mu sync.Mutex

	for i := 0; i < numTasks; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != numTasks {
		t.Errorf("Expected counter to be %d, got %d", numTasks, counter)
	}
}
