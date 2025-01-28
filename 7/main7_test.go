package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
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

	expected := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
	}

	for val := range merged {
		if !expected[val] {
			t.Errorf("Unexpected value: %d", val)
		}
		delete(expected, val)
	}

	if len(expected) != 0 {
		t.Errorf("Not all expected values were received: %v", expected)
	}
}
