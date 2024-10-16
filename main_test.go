package main

import (
	"testing"
)

// func TestProductExceptSelf(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		numbers  []int
// 		expected []int
// 	}{
// 		{
// 			name:     "Positive numbers",
// 			numbers:  []int{1, 2, 3, 4},
// 			expected: []int{24, 12, 8, 6},
// 		},
// 		{
// 			name:     "Mixed numbers",
// 			numbers:  []int{-1, 1, 0, -3, 3},
// 			expected: []int{0, 0, 9, 0, 0},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := productExceptSelf(tt.numbers)
// 			if !reflect.DeepEqual(got, tt.expected) {
// 				t.Errorf("sumArray(%v) = %v, want %v", tt.numbers, got, tt.expected)
// 			}
// 		})
// 	}
// }

// func TestMajorityElement(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		numbers  []int
// 		expected int
// 	}{
// 		{
// 			name:     "few number",
// 			numbers:  []int{3, 2, 3},
// 			expected: 3,
// 		},
// 		{
// 			name:     "edge number",
// 			numbers:  []int{6, 5, 5},
// 			expected: 5,
// 		},
// 		{
// 			name:     "more number",
// 			numbers:  []int{2, 2, 1, 1, 1, 2, 2},
// 			expected: 2,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := majorityElement(tt.numbers)
// 			if got != tt.expected {
// 				t.Errorf("majorityElement(%v) = %v, want %v", tt.numbers, got, tt.expected)
// 			}
// 		})
// 	}
// }

// func TestIncreasingTriplet(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		numbers  []int
// 		expected bool
// 	}{
// 		{
// 			name:     "few number",
// 			numbers:  []int{1, 2, 3, 4, 5},
// 			expected: true,
// 		},
// 		{
// 			name:     "edge number",
// 			numbers:  []int{5, 4, 3, 2, 1},
// 			expected: false,
// 		},
// 		{
// 			name:     "more number",
// 			numbers:  []int{2, 1, 5, 0, 4, 6},
// 			expected: true,
// 		},
// 		{
// 			name:     "edgy number",
// 			numbers:  []int{20, 100, 10, 12, 5, 13},
// 			expected: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := increasingTriplet(tt.numbers)
// 			if got != tt.expected {
// 				t.Errorf("majorityElement(%v) = %v, want %v", tt.numbers, got, tt.expected)
// 			}
// 		})
// 	}
// }

// func TestCompress(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		chars    []byte
// 		expected int
// 	}{
// 		{
// 			name:     "few number",
// 			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
// 			expected: 6,
// 		},
// 		{
// 			name:     "edge number",
// 			chars:    []byte{'a'},
// 			expected: 1,
// 		},
// 		{
// 			name:     "more number",
// 			chars:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
// 			expected: 4,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := compress(tt.chars)
// 			if got != tt.expected {
// 				t.Errorf("majorityElement(%v) = %v, want %v", tt.chars, got, tt.expected)
// 			}
// 		})
// 	}
// }
//

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "few number",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "edge number",
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.expected {
				t.Errorf("majorityElement(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.expected)
			}
		})
	}
}
