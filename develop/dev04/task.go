package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words *[]string) map[string][]string {
	// Создаем карту для хранения множеств анаграмм
	anagramSets := make(map[string][]string)

	// Создаем временную карту для проверки уникальности слов
	seen := make(map[string]bool)

	// Проходим по всем словам во входном массиве
	for _, word := range *words {
		// Приводим слово к нижнему регистру и сортируем его символы
		wordLower := strings.ToLower(word)
		wordSorted := sortString(wordLower)

		// Если слово уже было обработано как анаграмма, пропускаем его
		if _, found := seen[wordSorted]; found {
			continue
		}

		// Подготавливаем множество анаграмм для текущего слова
		anagramSet := []string{word}

		// Добавляем текущее слово во временную карту для проверки уникальности
		seen[wordSorted] = true

		// Проверяем остальные слова в массиве на анаграммы
		for _, otherWord := range *words {
			// Приводим слово к нижнему регистру и сортируем его символы
			otherWordLower := strings.ToLower(otherWord)
			otherWordSorted := sortString(otherWordLower)

			// Если это не текущее слово и оно является анаграммой, добавляем его в множество анаграмм
			if word != otherWord && wordSorted == otherWordSorted {
				anagramSet = append(anagramSet, otherWord)
				// Добавляем анаграмму во временную карту для проверки уникальности
				seen[otherWordSorted] = true
			}
		}

		// Если множество анаграмм имеет более одного элемента, добавляем его в результат
		if len(anagramSet) > 1 {
			sort.Strings(anagramSet)
			anagramSets[word] = anagramSet
		}
	}

	return anagramSets
}

// Вспомогательная функция для сортировки символов слова
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	words := []string{"пятак", "пятка", "кулон", "тяпка", "листок", "слиток", "клоун", "столик"}
	result := findAnagramSets(&words)
	// Вывод результата
	for key, value := range result {
		fmt.Println(key, ":", value)
	}
}
