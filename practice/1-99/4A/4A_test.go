package main

import (
	"testing"
)

func TestCF4A(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{8, "YES"},
		{11, "NO"},
		{2, "NO"},    // edge case
		{4, "YES"},   // smallest even number greater than 2
		{100, "YES"}, // large even number
	}

	for _, test := range tests {
		output := CF4A(test.input)
		if output != test.expected {
			t.Errorf("For input %d, expected %s, but got %s", test.input, test.expected, output)
		}
	}
}
