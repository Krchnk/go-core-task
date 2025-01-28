package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()
	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	merged := mergeChannels(ch1, ch2, ch3)
	for val := range merged {
		fmt.Println(val)
	}
}
