package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "Test 1",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "Test 2",
			slice1: []string{"a", "b", "c"},
			slice2: []string{"b", "c", "d"},
			want:   []string{"a"},
		},
		{
			name:   "Test 3",
			slice1: []string{"x", "y", "z"},
			slice2: []string{"a", "b", "c"},
			want:   []string{"x", "y", "z"},
		},
		{
			name:   "Test 4",
			slice1: []string{},
			slice2: []string{"a", "b", "c"},
			want:   []string{},
		},
		{
			name:   "Test 5",
			slice1: []string{"a", "b", "c"},
			slice2: []string{},
			want:   []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
