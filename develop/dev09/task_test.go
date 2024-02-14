package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	// Создание временного HTTP-сервера для тестирования
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Отправка тестового HTML-контента
		htmlContent := `<html><body><h1>Hello, World!</h1></body></html>`
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(htmlContent))
	}))
	defer testServer.Close()

	// Создание временного файла для сохранения
	tempFileName := "test.html"
	defer os.Remove(tempFileName)

	// Выполнение тестируемой функции
	err := downloadPage(testServer.URL, tempFileName)
	if err != nil {
		t.Fatalf("ошибка при загрузке страницы: %v", err)
	}

	// Чтение содержимого файла
	content, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		t.Fatalf("ошибка при чтении файла: %v", err)
	}

	// Проверка содержимого файла
	expectedContent := []byte(`<html><body><h1>Hello, World!</h1></body></html>`)
	if !bytes.Equal(content, expectedContent) {
		t.Error("ожидается содержимое файла:", string(expectedContent))
		t.Error("фактическое содержимое файла:", string(content))
	}
}

func TestMain(m *testing.M) {
	// Запуск тестов
	exitCode := m.Run()

	// Очистка временных файлов и завершение теста
	os.Remove("test.html")
	os.Exit(exitCode)
}
