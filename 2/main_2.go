package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRandomSlice() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = r.Intn(100)
	}
	return slice
}

func sliceExample(slice []int) []int {
	var evenSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	originalSlice := createRandomSlice()
	fmt.Println("Original Slice:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Numbers Slice:", evenSlice)

	addedSlice := addElements(originalSlice, 100)
	fmt.Println("Slice with Added Element:", addedSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	removedSlice := removeElement(originalSlice, 3)
	fmt.Println("Slice with Removed Element:", removedSlice)
}
