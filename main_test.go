package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{
			name:     "Positive numbers",
			numbers:  []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Mixed numbers",
			numbers:  []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.numbers)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("sumArray(%v) = %v, want %v", tt.numbers, got, tt.expected)
			}
		})
	}
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "few number",
			numbers:  []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "edge number",
			numbers:  []int{6, 5, 5},
			expected: 5,
		},
		{
			name:     "more number",
			numbers:  []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.numbers)
			if got != tt.expected {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.numbers, got, tt.expected)
			}
		})
	}
}

func TestRotate(nums []int, k int) {
	tests := []struct {
		name    string
		numbers []int
		k       int
	}{
		{
			name:    "few number",
			numbers: []int{1, 2, 3, 4, 5, 6, 7},
			k:       3,
		},
		{
			name:    "edge number",
			numbers: []int{-1, -100, 3, 99},
			k:       5,
		},
	}
}
