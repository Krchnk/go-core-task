package main

import (
	"math/rand"
	"time"
)

func RandomGenerator(done <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			select {
			case <-done:
				return
			case out <- localRand.Intn(100):
			}
		}
	}()

	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	randomChan := RandomGenerator(done)
	for i := 0; i < 10; i++ {
		println(<-randomChan)
	}
}
