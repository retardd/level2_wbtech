package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		input    string
		expected string
	}{
		{
			name:     "Cut_first_field_with_default_delimiter",
			args:     []string{"-f", "1"},
			input:    "hello\tworld\nfoo\tbar\n",
			expected: "hello\nfoo\n",
		},
		{
			name:     "Cut_second_field_with_custom_delimiter",
			args:     []string{"-f", "2", "-d", ";"},
			input:    "apple;orange;banana\none;two;three\n",
			expected: "orange\ntwo\n",
		},
		{
			name:     "Cut_third_field_only_for_lines_with_delimiter",
			args:     []string{"-f", "3", "-s"},
			input:    "City\nNew York\nLondon\n",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "task.go")
			cmd.Args = append(cmd.Args, tt.args...)

			// Передача входных данных построчно
			cmd.Stdin = strings.NewReader(tt.input)

			var out bytes.Buffer
			cmd.Stdout = &out

			err := cmd.Run()
			if err != nil {
				t.Fatalf("error executing command: %v", err)
			}

			got := out.String()
			if got != tt.expected {
				t.Errorf("unexpected output:\n%s\nexpected:\n%s\n", got, tt.expected)
			}
		})
	}
}
