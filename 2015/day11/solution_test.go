package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test1_1",
			input:    "abcdefgh",
			expected: "abcdffaa",
		},
		{
			name:     "Test1_2",
			input:    "ghijklmn",
			expected: "ghjaabcc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve1(tt.input)
			if err != nil {
				t.Fatalf("solve1() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve1()  = %s;  expected %s", got, tt.expected)
			}
		})
	}
}

func TestSolve2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test2_1",
			input:    "abcdffaa",
			expected: "abcdffbb",
		},
		{
			name:     "Test2_2",
			input:    "ghjaabcc",
			expected: "ghjbbcdd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.input)
			if err != nil {
				t.Fatalf("solve2() returned an unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("solve2() = %s; expected %s", got, tt.expected)
			}
		})
	}
}
