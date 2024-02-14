package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestShell(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expectedOut string
	}{
		{
			name:        "TestCDPWD",
			commands:    "cd pwd\n\\quit\n",
			expectedOut: "",
		},
		{
			name:        "TestPWD",
			commands:    "pwd\n\\quit\n",
			expectedOut: "$ D:\\GITHUB\\GOLANG-WB\\l2\\develop\\dev08\n$ ",
		},
		{
			name:        "TestCDnull",
			commands:    "cd\n\\quit\n",
			expectedOut: "Необходимо указать путь для cd",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Error creating pipe: %v", err)
			}
			original := os.Stdout
			os.Stdout = w

			go func() {
				defer w.Close()
				cmd := exec.Command("go", "run", "task.go")
				cmd.Stdin = strings.NewReader(test.commands)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			}()

			var buf bytes.Buffer
			io.Copy(&buf, r)
			os.Stdout = original

			got := buf.String()
			if !strings.Contains(got, test.expectedOut) {
				t.Errorf("Unexpected output: got %q, want %q", got, test.expectedOut)
			}
		})
	}
}
