package main

import "testing"

func TestSolve1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "nice 1",
			input:    "ugknbfddgicrmopn",
			expected: 1,
		},
		{
			name:     "nice 2",
			input:    "aaa",
			expected: 1,
		},
		{
			name:     "naughty 1",
			input:    "jchzalrnumimnmhp",
			expected: 0,
		},
		{
			name:     "naughty 2",
			input:    "haegwjzuvuyypxyu",
			expected: 0,
		},
		{
			name:     "naughty 3",
			input:    "dvszwmarrgswjxmb",
			expected: 0,
		},
		{
			name:     "Empty String",
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
			name:     "nice 1",
			input:    "qjhvhtzxzqqjkmpb",
			expected: 1,
		},
		{
			name:     "nice 2",
			input:    "xxyxx",
			expected: 1,
		},
		{
			name:     "naughty 1",
			input:    "uurcxstgmygtbstg",
			expected: 0,
		},
		{
			name:     "naughty 2",
			input:    "ieodomkazucvgmuy",
			expected: 0,
		},
		{
			name:     "empty string",
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
				t.Errorf("solve2() = %d; exected %d", got, tt.expected)
			}
		})
	}
}
