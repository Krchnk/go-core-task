package main

import (
	"math"
	"testing"
)

func TestProcessNumbers(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go processNumbers(input, output)

	go func() {
		input <- 2
		input <- 3
		close(input)
	}()

	expected := []float64{math.Pow(2, 3), math.Pow(3, 3)}
	i := 0
	for result := range output {
		if result != expected[i] {
			t.Errorf("Expected %f, got %f", expected[i], result)
		}
		i++
	}
}
