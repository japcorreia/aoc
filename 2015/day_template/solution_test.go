package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "",
			input:    "",
			expected: 0,
		},
		{
			name:     "",
			input:    "",
			expected: 0,
		},
		{
			name:     "",
			input:    "",
			expected: 0,
		},
		{
			name:     "Empty Path",
			input:    "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve1(tt.input)
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
			name:     "",
			input:    "",
			expected: 0,
		},
		{
			name:     "",
			input:    "",
			expected: 0,
		},
		{
			name:     "",
			input:    "",
			expected: 0,
		},
		{
			name:     "empty path",
			input:    "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.input)
			if err != nil {
				t.Fatalf("solve2() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve2() = %d; expected %d", got, tt.expected)
			}
		})
	}
}
