package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "      ",
			expected: []string{},
		},
		{
			input:    "Hello, World!",
			expected: []string{"hello,","world!"},
		},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(tc.expected) {
			t.Errorf("expected length: %v, actual length: %v", len(tc.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := tc.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("expected: %v, actual: %v", expectedWord, word)
			}
		}
	}
}
