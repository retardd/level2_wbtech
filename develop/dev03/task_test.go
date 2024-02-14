package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	// Define test cases
	testCases := []struct {
		args        string
		expectedOut string
		input       string
	}{
		{"", "aaa\naaa\nbbb\nbbb\nccc\nccc\n", "bbb\nccc\naaa\nccc\naaa\nbbb\n"},
		{"-n", "1\n3\n3\n4\n6\n7\n", "3\n1\n3\n4\n7\n6\n"},
		{"-r", "ccc\nccc\nbbb\nbbb\naaa\naaa\n", "bbb\nccc\naaa\nccc\naaa\nbbb\n"},
		{"-u", "aaa\nbbb\nccc\n", "bbb\nccc\naaa\nccc\naaa\nbbb\n"},
		{"-n -r", "7\n6\n4\n3\n3\n1\n", "3\n1\n3\n4\n7\n6\n"},
		{"-r -u", "ccc\nbbb\naaa\n", "bbb\nccc\naaa\nccc\naaa\nbbb\n"},
		{"-n -r -u", "7\n6\n4\n3\n1\n", "3\n1\n3\n4\n7\n6\n"},
	}

	// Save and restore os.Stdout
	old := os.Stdout
	defer func() { os.Stdout = old }()

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.args, func(t *testing.T) {
			// Create a pipe to capture output
			r, w, _ := os.Pipe()
			os.Stdout = w

			tmpfile, err := ioutil.TempFile("", "testfile.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			testData := tc.input
			if _, err := tmpfile.Write([]byte(testData)); err != nil {
				tmpfile.Close()
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}
			// Run the sort command with the specified arguments
			oldArgs := os.Args
			os.Args = []string{"cmd", tmpfile.Name()}
			os.Args = append(os.Args, strings.Fields(tc.args)...)
			main()
			os.Args = oldArgs

			// Read output from the pipe
			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)

			// Check the output
			got := buf.String()
			if got != tc.expectedOut {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tc.expectedOut, got)
			}
		})
	}
}
