package main

import (
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	randomChan := RandomGenerator(done)

	for i := 0; i < 10; i++ {
		select {
		case num := <-randomChan:
			if num < 0 || num >= 100 {
				t.Errorf("Ожидалось число в диапазоне [0, 100), получено: %d", num)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("Таймаут при чтении из канала")
		}
	}
}

func TestRandomGenerator_Stop(t *testing.T) {
	done := make(chan struct{})
	randomChan := RandomGenerator(done)

	close(done)

	time.Sleep(10 * time.Millisecond)

	select {
	case _, ok := <-randomChan:
		if ok {
			t.Error("Ожидалось, что канал будет закрыт")
		}
	default:

	}
}
