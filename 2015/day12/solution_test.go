package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test1_1",
			input:    "[1,2,3]",
			expected: 6,
		},
		{
			name:     "Test1_2",
			input:    "[[[3]]]",
			expected: 3,
		},
		{
			name:     "Test1_3",
			input:    "[-1,{\"a\":1}]",
			expected: 0,
		},
		{
			name:     "Test1_4",
			input:    "{\"a\":[]}",
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
			name:     "Test 2_1",
			input:    "[1,2,3]",
			expected: 6,
		},
		{
			name:     "Test 2_2",
			input:    "[1,{\"c\":\"red\",\"b\":2},3]",
			expected: 4,
		},
		{
			name:     "Test 2_3",
			input:    "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}",
			expected: 0,
		},
		{
			name:     "Test 2_4",
			input:    "{\"a\":{\"b\":1},\"c\":\"red\"}",
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
