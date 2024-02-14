package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

//Тесты почему-то работают адекватно только если через возможности IDE запускать по одному:(

func TestGrep(t *testing.T) {
	// Создаем временный файл для тестов
	tempFile := "test.txt"
	file, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer file.Close()

	// Записываем строки в файл
	lines := []string{
		"Hello world",
		"Test line 1",
		"Test line 2",
		"Another line",
		"Test line 3",
	}
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	// Сохраняем оригинальные потоки вывода
	originalStdout := os.Stdout
	originalStderr := os.Stderr

	// Создаем каналы для хранения оригинальных потоков вывода
	rStdout, wStdout, _ := os.Pipe()
	rStderr, wStderr, _ := os.Pipe()

	// Перенаправляем stdout и stderr на буферы
	os.Stdout = wStdout
	os.Stderr = wStderr

	// Создаем канал для получения результатов вывода
	mergedOutput := make(chan string)

	// Восстанавливаем оригинальные потоки вывода после завершения
	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	// Тесты
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Basic grep",
			args:     []string{"-n", "Test", tempFile},
			expected: "2:Test line 1\n3:Test line 2\n5:Test line 3\n",
		},
		{
			name:     "Case insensitive grep",
			args:     []string{"-n", "-i", "hello", tempFile},
			expected: "1:hello world\n",
		},
		{
			name:     "Print count",
			args:     []string{"-c", "Test", tempFile},
			expected: "3\n",
		},
		{
			name:     "Print before",
			args:     []string{"-B", "1", "line", tempFile},
			expected: "Hello world\nTest line 1\nTest line 2\nAnother line\nTest line 3\n",
		},
		{
			name:     "Print after",
			args:     []string{"-a", "1", "line", tempFile},
			expected: "Test line 1\nTest line 2\nAnother line\nTest line 3\n",
		},
		{
			name:     "Print context",
			args:     []string{"-C", "1", "Another line", tempFile},
			expected: "Another line\nTest line 2\nTest line 3\n",
		},
		{
			name:     "Invert match",
			args:     []string{"-v", "-n", "line", tempFile},
			expected: "1:Hello world\n",
		},
		{
			name:     "Fixed strings",
			args:     []string{"-F", "-n", "Test line 1", tempFile},
			expected: "2:Test line 1\n",
		},
	}

	// Запуск тестов
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Передаем тестовые аргументы
			os.Args = append([]string{"./grep"}, tc.args...)

			var wg sync.WaitGroup

			wg.Add(1)
			go func() {
				main()
				wStdout.Close()
				wStderr.Close()
				wg.Done()
			}()

			// Считываем вывод из stdout и stderr
			var stdoutBuf, stderrBuf bytes.Buffer
			wg.Add(1)
			go func() {
				io.Copy(&stdoutBuf, rStdout)
				mergedOutput <- stdoutBuf.String()
				wg.Done()
			}()
			wg.Add(1)
			go func() {
				io.Copy(&stderrBuf, rStderr)
				mergedOutput <- stderrBuf.String()
				wg.Done()
			}()

			// Получаем результаты вывода
			var output string
			for i := 0; i < 2; i++ {
				output += <-mergedOutput
			}
			wg.Wait()

			// Проверяем результаты
			if output != tc.expected {
				t.Errorf("Expected output: %s, but got: %s", tc.expected, output)
			}
		})
	}
}
