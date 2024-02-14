package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

func main() {
	var (
		after      int
		before     int
		context    int
		count      bool
		ignoreCase bool
		invert     bool
		fixed      bool
		lineNum    bool
	)
	// Определение флагов
	pflag.IntVarP(&after, "after", "a", 0, "Print +N lines after each match") // Используйте pflag.IntVarP
	pflag.IntVarP(&before, "before", "B", 0, "Print +N lines before each match")
	pflag.IntVarP(&context, "context", "C", 0, "Print ±N lines around each match")
	pflag.BoolVarP(&count, "count", "c", false, "Print only a count of selected lines per FILE")
	pflag.BoolVarP(&ignoreCase, "ignore-case", "i", false, "Ignore case distinctions")
	pflag.BoolVarP(&invert, "invert", "v", false, "Invert the sense of matching")
	pflag.BoolVarP(&fixed, "fixed", "F", false, "Interpret PATTERN as a list of fixed strings")
	pflag.BoolVarP(&lineNum, "line-num", "n", false, "Print line number with output lines")
	pflag.Parse()

	// Получение паттерна для поиска
	args := pflag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		pflag.PrintDefaults()
		os.Exit(1)
	}

	pattern := args[0]

	// Открытие файла или стандартного ввода
	var input *os.File
	var err error
	if pflag.NArg() == 2 {
		input, err = os.Open(pflag.Arg(1))
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer input.Close()
	} else {
		fmt.Println("pflag.NArg()", err)
		os.Exit(1)
	}

	// Чтение из входного потока
	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var matches int
	var buffer []string
	lineNumber := 0
	for lineNumber < len(lines) {
		line := lines[lineNumber]
		if ignoreCase {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}
		match := false
		if fixed {
			match = strings.Contains(line, pattern)
		} else {
			match = strings.Contains(line, pattern)
		}
		if len(buffer) > 0 {
			for _, b := range buffer {
				if lineNum {
					fmt.Printf("%d:", lineNumber+1)
				}
				fmt.Println(b)
			}
			buffer = nil
		}
		if match != invert {
			if count {
				matches++
				lineNumber++
			} else {
				if lineNum {
					fmt.Printf("%d:", lineNumber+1)
				}
				fmt.Println(line)
				lineNumber++
				if after > 0 {
					for i := 1; i <= after && lineNumber < len(lines); i++ {
						line = lines[lineNumber]
						if lineNum {
							fmt.Printf("%d:", lineNumber+1)
						}
						fmt.Println(line)
						lineNumber++
					}
				}
				if context > 0 {
					j := 0
					for j < context && lineNumber-2-j >= 0 {
						buffer = append(buffer, lines[lineNumber-2-j])
						j++
					}
					j = 0
					for j < context && lineNumber+j < len(lines) {
						buffer = append(buffer, lines[lineNumber+j])
						j++
					}
					for len(buffer) > context*2 {
						buffer = buffer[1:]
					}
				}
			}
		} else if before > 0 {
			if len(buffer) == before {
				buffer = buffer[1:]
			}
			buffer = append(buffer, line)
			lineNumber++
		} else {
			lineNumber++
		}
	}

	// Вывод количества совпадений
	if count {
		fmt.Println(matches)
	}
}
