package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(slice)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}
	result := addElements(slice, 4)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3}
	result := copySlice(slice)

	if len(result) != len(slice) {
		t.Errorf("Expected %v, got %v", slice, result)
	}

	for i := range result {
		if result[i] != slice[i] {
			t.Errorf("Expected %v, got %v", slice, result)
		}
	}

	slice[0] = 100
	if result[0] == slice[0] {
		t.Errorf("Changes in original slice affected the copy")
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}
	result := removeElement(slice, 2)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}
