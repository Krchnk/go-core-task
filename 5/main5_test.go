package main

import (
	"testing"
)

func TestFindIntersections(t *testing.T) {
	tests := []struct {
		name              string
		a, b              []int
		wantBool          bool
		wantIntersections []int
	}{
		{
			name:              "No intersections",
			a:                 []int{1, 2, 3},
			b:                 []int{4, 5, 6},
			wantBool:          false,
			wantIntersections: []int{},
		},
		{
			name:              "Single intersection",
			a:                 []int{1, 2, 3},
			b:                 []int{3, 4, 5},
			wantBool:          true,
			wantIntersections: []int{3},
		},
		{
			name:              "Multiple intersections",
			a:                 []int{65, 3, 58, 678, 64},
			b:                 []int{64, 2, 3, 43},
			wantBool:          true,
			wantIntersections: []int{64, 3},
		},
		{
			name:              "Empty slices",
			a:                 []int{},
			b:                 []int{},
			wantBool:          false,
			wantIntersections: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotIntersections := FindIntersections(tt.a, tt.b)
			if gotBool != tt.wantBool {
				t.Errorf("FindIntersections() gotBool = %v, want %v", gotBool, tt.wantBool)
			}
			if !equalSlices(gotIntersections, tt.wantIntersections) {
				t.Errorf("FindIntersections() gotIntersections = %v, want %v", gotIntersections, tt.wantIntersections)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
