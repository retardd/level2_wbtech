package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// Установка флага таймаута
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	// Получение аргументов командной строки
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Использование: go-telnet [--timeout=10s] host port")
		return
	}
	host := args[0]
	port := args[1]

	// Установка соединения
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Успешно подключен к %s:%s\n", host, port)

	// Канал для ожидания сигнала завершения
	done := make(chan struct{})

	// Запуск горутины для чтения из сокета и вывода в STDOUT
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			log.Printf("Ошибка при чтении из сокета: %v", err)
		}
		close(done)
	}()

	// Запуск горутины для чтения из STDIN и записи в сокет
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			log.Printf("Ошибка при записи в сокет: %v", err)
		}
		close(done)
	}()

	// Ожидание сигнала завершения (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	// Закрытие соединения
	conn.Close()

	// Ожидание завершения горутин
	<-done
}
