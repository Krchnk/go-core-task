package main

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestPrintType(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}

	for _, test := range tests {
		result := printType(test.input)
		if result != test.expected {
			t.Errorf("printType(%v) = %v, ожидалось %v", test.input, result, test.expected)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "42"},
		{3.14, "3.14"},
		{"Golang", "Golang"},
		{true, "true"},
		{complex64(1 + 2i), "(1+2i)"},
	}

	for _, test := range tests {
		result := toString(test.input)
		if result != test.expected {
			t.Errorf("toString(%v) = %v, ожидалось %v", test.input, result, test.expected)
		}
	}
}

func TestConcatenateVariables(t *testing.T) {
	result := concatenateVariables(42, 3.14, "Golang", true, complex64(1+2i))
	expected := "42 3.14 Golang true (1+2i) "
	if result != expected {
		t.Errorf("concatenateVariables() = %v, ожидалось %v", result, expected)
	}
}

func TestStringToRunes(t *testing.T) {
	input := "Golang"
	expected := []rune{'G', 'o', 'l', 'a', 'n', 'g'}
	result := stringToRunes(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("stringToRunes(%v) = %v, ожидалось %v", input, result, expected)
	}
}

func TestHashRunes(t *testing.T) {
	runes := []rune{'G', 'o', 'l', 'a', 'n', 'g'}
	salt := "go-2024"
	expected := sha256.Sum256([]byte("Golgo-2024ang"))
	expectedHash := hex.EncodeToString(expected[:])
	result := hashRunes(runes, salt)
	if result != expectedHash {
		t.Errorf("hashRunes() = %v, ожидалось %v", result, expectedHash)
	}
}
