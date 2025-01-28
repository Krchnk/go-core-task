package main

import (
	"fmt"
)

func FindIntersections(a, b []int) (bool, []int) {

	elements := make(map[int]bool)
	for _, v := range a {
		elements[v] = true
	}

	var intersections []int

	for _, v := range b {
		if elements[v] {
			intersections = append(intersections, v)
			delete(elements, v)
		}
	}

	if len(intersections) > 0 {
		return true, intersections
	}

	return false, intersections
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersections := FindIntersections(a, b)
	fmt.Println(hasIntersection, intersections)
}
