package main

import (
	"bytes"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func startTestServer() {
	// Эмуляция сервера. Здесь сервер создается только для теста и не выполняется на реальном порту.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":8080", nil)
}

func TestTelnetClient(t *testing.T) {
	// Запуск временного HTTP-сервера
	go startTestServer()

	// Запуск вашего telnet-клиента через go run task.go в отдельном процессе
	cmd := exec.Command("go", "run", "task.go", "localhost", "8080")

	// Подготовка запроса GET
	getRequest := []byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n")

	// Запись запроса в stdin telnet-клиента
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("ошибка при создании входного канала: %v", err)
	}
	defer stdin.Close()

	if _, err := stdin.Write(getRequest); err != nil {
		t.Fatalf("ошибка при записи в stdin: %v", err)
	}

	// Добавляем задержку перед отправкой второго Enter
	time.Sleep(100 * time.Millisecond)

	// Отправляем второй Enter
	if _, err := stdin.Write([]byte("\r\n")); err != nil {
		t.Fatalf("ошибка при записи в stdin: %v", err)
	}

	// Создание буфера для чтения ответа
	var buf bytes.Buffer
	cmd.Stdout = &buf

	// Запуск процесса и ожидание завершения
	if err := cmd.Start(); err != nil {
		t.Fatalf("ошибка при запуске процесса: %v", err)
	}
	defer cmd.Process.Kill()

	time.Sleep(4000 * time.Millisecond)

	// Проверка полученного ответа
	expectedResponse := "Hello, World!"
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	actualResponse := strings.TrimSpace(lines[len(lines)-1])
	if actualResponse != expectedResponse {
		t.Fatalf("ожидается ответ %q, получено %q", expectedResponse, actualResponse)
	}
}
