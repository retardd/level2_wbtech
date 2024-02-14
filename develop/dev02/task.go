package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func unpackString(s string) (string, error) {
	var result string
	var escape bool

	for i := 0; i < len(s); i++ {
		if escape {
			escape = false
			result += string(s[i])
			continue
		}

		if s[i] == '\\' {
			escape = true
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			return "", errors.New("invalid string")
		}

		result += string(s[i])

		if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
			count, _ := strconv.Atoi(string(s[i+1]))
			result += strings.Repeat(string(s[i]), count-1)
			i++
		}
	}

	return result, nil
}

func main() {
	inputs := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}

	for _, input := range inputs {
		unpacked, err := unpackString(input)
		if err != nil {
			fmt.Printf("Ошибка для строки '%s': %s\n", input, err)
		} else {
			fmt.Printf("Распакованная строка для '%s': %s\n", input, unpacked)
		}
	}
}
