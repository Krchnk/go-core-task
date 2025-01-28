package main

import (
	"fmt"
	"math"
)

func processNumbers(input <-chan uint8, output chan<- float64) {
	for num := range input {
		result := math.Pow(float64(num), 3)
		output <- result
	}
	close(output)
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go processNumbers(input, output)

	go func() {
		numbers := []uint8{1, 2, 3, 4, 5}
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	for result := range output {
		fmt.Println(result)
	}
}
