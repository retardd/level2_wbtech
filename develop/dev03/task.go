package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/spf13/cobra"
)

type sortBy int

const (
	byDefault sortBy = iota
	byNumeric
	byMonth
)

type sortMode struct {
	column           int
	by               sortBy
	reverse          bool
	unique           bool
	ignoreWhitespace bool
	checkSorted      bool
	byHuman          bool
	byNumeric        bool
	byMonth          bool
}

func main() {
	var mode sortMode

	var rootCmd = &cobra.Command{
		Use:   "sort",
		Short: "Utility to sort lines in a file",
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			lines, err := readLines(fileName)
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}

			if mode.checkSorted {
				if isSorted(lines, mode) {
					fmt.Println("File is sorted")
				} else {
					fmt.Println("File is not sorted")
				}
				return
			}

			lines = sortLines(lines, mode)

			for _, line := range lines {
				fmt.Println(line)
			}
		},
	}

	rootCmd.Flags().IntVarP(&mode.column, "key", "k", -1, "column number for sorting")
	rootCmd.Flags().BoolVarP(&mode.reverse, "reverse", "r", false, "sort in reverse order")
	rootCmd.Flags().BoolVarP(&mode.unique, "unique", "u", false, "suppress repeated lines")
	rootCmd.Flags().BoolVarP(&mode.ignoreWhitespace, "ignore-leading-blanks", "b", false, "ignore leading whitespace")
	rootCmd.Flags().BoolVarP(&mode.checkSorted, "check-sorted", "c", false, "check if file is sorted")
	rootCmd.Flags().BoolVarP(&mode.byHuman, "human-numeric-sort", "H", false, "sort human-readable numbers")
	rootCmd.Flags().BoolVarP(&mode.byNumeric, "numeric-sort", "n", false, "sort numerically")
	rootCmd.Flags().BoolVarP(&mode.byMonth, "month-sort", "M", false, "sort by month")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		var err = file.Close()
		if err != nil {

		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isSorted(lines []string, mode sortMode) bool {
	var comparator func(i, j int) bool
	switch {
	case mode.byNumeric:
		comparator = func(i, j int) bool {
			num1, _ := strconv.Atoi(strings.Fields(lines[i])[mode.column-1])
			num2, _ := strconv.Atoi(strings.Fields(lines[j])[mode.column-1])
			if mode.reverse {
				return num1 >= num2
			}
			return num1 <= num2
		}
	case mode.byMonth:
		comparator = func(i, j int) bool {
			date1, _ := time.Parse("Jan", strings.Fields(lines[i])[mode.column-1])
			date2, _ := time.Parse("Jan", strings.Fields(lines[j])[mode.column-1])
			if mode.reverse {
				return date1.After(date2)
			}
			return date1.Before(date2)
		}
	default:
		comparator = func(i, j int) bool {
			if mode.reverse {
				return lines[i] >= lines[j]
			}
			return lines[i] <= lines[j]
		}
	}

	return sort.SliceIsSorted(lines, comparator)
}

func sortLines(lines []string, mode sortMode) []string {
	var lessFunc func(i, j int) bool
	switch {
	case mode.byMonth:
		if mode.column < 0 {
			// сортировка по всему тексту
			lessFunc = func(i, j int) bool {
				date1, _ := time.Parse("Jan", strings.Fields(lines[i])[0])
				date2, _ := time.Parse("Jan", strings.Fields(lines[j])[0])
				if mode.reverse {
					return date1.After(date2)
				}
				return date1.Before(date2)
			}
		} else {
			// сортировка по указанной колонке
			lessFunc = func(i, j int) bool {
				date1, _ := time.Parse("Jan", strings.Fields(lines[i])[mode.column-1])
				date2, _ := time.Parse("Jan", strings.Fields(lines[j])[mode.column-1])
				if mode.reverse {
					return date1.After(date2)
				}
				return date1.Before(date2)
			}
		}
	case mode.byNumeric:
		if mode.column < 0 {
			// сортировка по всему тексту
			lessFunc = func(i, j int) bool {
				num1, _ := strconv.Atoi(strings.Fields(lines[i])[0])
				num2, _ := strconv.Atoi(strings.Fields(lines[j])[0])
				if mode.reverse {
					return num1 >= num2
				}
				return num1 <= num2
			}
		} else {
			// сортировка по указанной колонке
			lessFunc = func(i, j int) bool {
				num1, _ := strconv.Atoi(strings.Fields(lines[i])[mode.column-1])
				num2, _ := strconv.Atoi(strings.Fields(lines[j])[mode.column-1])
				if mode.reverse {
					return num1 >= num2
				}
				return num1 <= num2
			}
		}
	default: // сортировка по умолчанию
		if mode.column < 0 {
			// сортировка по всему тексту
			lessFunc = func(i, j int) bool {
				str1 := lines[i]
				str2 := lines[j]
				if mode.ignoreWhitespace {
					str1 = strings.TrimSpace(str1) // игнорировать конечные пробелы
					str2 = strings.TrimSpace(str2)
				}
				if mode.reverse {
					return str1 >= str2
				}
				return str1 <= str2
			}
		} else {
			// сортировка по указанной колонке
			lessFunc = func(i, j int) bool {
				columns1 := strings.Fields(lines[i])
				columns2 := strings.Fields(lines[j])
				if len(columns1) <= mode.column-1 {
					return false
				}
				if len(columns2) <= mode.column-1 {
					return true
				}
				str1 := columns1[mode.column-1]
				str2 := columns2[mode.column-1]
				if mode.ignoreWhitespace {
					str1 = strings.TrimSpace(str1) // игнорировать конечные пробелы
					str2 = strings.TrimSpace(str2)
				}
				if mode.reverse {
					return str1 >= str2
				}
				return str1 <= str2
			}
		}
	}

	sort.SliceStable(lines, lessFunc)

	if mode.unique {
		uniqueLines := make(map[string]bool)
		var result []string
		for _, line := range lines {
			if !uniqueLines[line] {
				uniqueLines[line] = true
				result = append(result, line)
			}
		}
		lines = result
	}

	return lines
}

func humanReadableNumber(str string) float64 {
	num, _ := strconv.ParseFloat(strings.TrimRightFunc(str, func(r rune) bool {
		return !unicode.IsDigit(r) && r != '.'
	}), 64)
	return num
}
