package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test 1",
			input:    "turn on 0,0 through 999,999\ntoggle 0,0 through 999,0\nturn off 499,499 through 500,500",
			expected: 998996,
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
			name:     "Test1",
			input:    "turn on 0,0 through 0,0",
			expected: 1,
		},
		{
			name:     "Test2",
			input:    "toggle 0,0 through 999,999",
			expected: 2000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.input)
			if err != nil {
				t.Fatalf("solve2() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve2() = %d; exected %d", got, tt.expected)
			}
		})
	}
}
