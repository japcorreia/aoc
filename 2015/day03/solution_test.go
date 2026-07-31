package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Short Delivery",
			input:    ">",
			expected: 2,
		},
		{
			name:     "Square Delivery",
			input:    "^>v<",
			expected: 4,
		},
		{
			name:     "Multiple Presents Delivery",
			input:    "^v^v^v^v^v",
			expected: 2,
		},
		{
			name:     "Empty Path",
			input:    "",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve1([]byte(tt.input))
			if err != nil {
				t.Fatalf("solve1() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve1()  = %d;  expected %d", got, tt.expected)
			}
		})
	}
}

func TestSolve2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Short Delivery",
			input:    "^v",
			expected: 3,
		},
		{
			name:     "Square Delivery",
			input:    "^>v<",
			expected: 3,
		},
		{
			name:     "Multiple Presents Delivery",
			input:    "^v^v^v^v^v",
			expected: 11,
		},
		{
			name:     "empty path",
			input:    "",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2([]byte(tt.input))
			if err != nil {
				t.Fatalf("solve2() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve2() = %d; exected %d", got, tt.expected)
			}
		})
	}
}
