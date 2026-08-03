package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint16
	}{
		{
			name:     "Test1",
			input:    "123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i",
			expected: 65412,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve1(tt.input, "h")
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
		expected uint16
	}{
		{
			name:     "Test 2",
			input:    "123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i",
			expected: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.input, "i", "y", 65412)
			if err != nil {
				t.Fatalf("solve2() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve2() = %d; expected %d", got, tt.expected)
			}
		})
	}
}
