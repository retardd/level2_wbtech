package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
	}

	for _, test := range tests {
		result, err := unpackString(test.input)
		if err != nil && !test.err {
			t.Errorf("unexpected error for input %q: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("unexpected result for input %q: got %q, expected %q", test.input, result, test.expected)
		}
	}
}
