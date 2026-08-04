package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		times    int
		expected int
	}{
		{
			name:     "Test 1_1",
			input:    "1",
			times:    1,
			expected: 2,
		},
		{
			name:     "Test 1_2",
			input:    "11",
			times:    1,
			expected: 2,
		},
		{
			name:     "Test 1_3",
			input:    "1211",
			times:    1,
			expected: 6,
		},
		{
			name:     "Test 1_4",
			input:    "111221",
			times:    1,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve1(tt.input, tt.times)
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
		times    int
		expected int
	}{
		{
			name:     "Test 2_1",
			input:    "234567",
			times:    20,
			expected: 2844,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.input, tt.times)
			if err != nil {
				t.Fatalf("solve2() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve2() = %d; expected %d", got, tt.expected)
			}
		})
	}
}
