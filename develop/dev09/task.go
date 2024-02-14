package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadPage(URL string, fileName string) error {
	// Выполнение HTTP-запроса
	response, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении HTTP-запроса: %v", err)
	}
	defer response.Body.Close()

	// Создание файла для сохранения
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer file.Close()

	// Копирование данных из ответа в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("ошибка при копировании данных: %v", err)
	}

	return nil
}

func main() {
	// Установка флагов командной строки
	fileNameFlag := flag.String("O", "index.html", "Имя файла для сохранения")
	flag.Parse()

	// Получение URL из аргументов командной строки
	URL := flag.Arg(0)
	if URL == "" {
		fmt.Println("Использование: ./wget <URL>")
		return
	}

	// Вызов функции загрузки страницы
	err := downloadPage(URL, *fileNameFlag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Страница успешно скачана и сохранена в файле: %s\n", *fileNameFlag)
}
