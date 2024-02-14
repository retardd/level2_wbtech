package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ") // Приглашение командной строки
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "\\quit" { // Выход из шелла
			break
		}

		commands := strings.Split(input, "|") // Разделение на команды по пайпу
		var output io.Reader

		for _, cmdStr := range commands {
			cmdStr = strings.TrimSpace(cmdStr)
			args := strings.Fields(cmdStr) // Разделение аргументов команды

			switch args[0] {
			case "cd":
				if len(args) < 2 {
					fmt.Println("Необходимо указать путь для cd")
					continue
				}
				if args[1] == "pwd" {
					args[1], err = os.Getwd()
				}
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println("Ошибка при смене директории:", err)
				}
			case "pwd":
				dir, err := os.Getwd()
				if err != nil {
					fmt.Println("Ошибка при получении текущей директории:", err)
				} else {
					fmt.Println(dir)
				}
			case "ps":
				cmd := exec.Command("ps")
				output, err := cmd.Output()
				if err != nil {
					fmt.Println("Ошибка при выполнении команды ps:", err)
				}
				fmt.Print(string(output))
			default:
				cmd := exec.Command(args[0], args[1:]...)
				if output != nil {
					cmd.Stdin = output
				}
				cmdOutput, err := cmd.Output()
				if err != nil {
					fmt.Println("Ошибка при выполнении команды:", err)
					break
				}
				output = strings.NewReader(string(cmdOutput))
			}
		}

		// Вывод результатов последней команды
		if output != nil {
			scanner := bufio.NewScanner(output)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка при чтении вывода команды:", err)
			}
		}
	}

	// Закрытие стандартного ввода после обработки всех команд
	os.Stdin.Close()
}
