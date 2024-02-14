package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	testCases := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		// Добавьте здесь другие тестовые случаи при необходимости
	}

	for _, tc := range testCases {
		result := findAnagramSets(&tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
		}
	}
}
