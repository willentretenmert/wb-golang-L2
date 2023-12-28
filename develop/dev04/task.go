package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(findAnagrams([]string{"пятак", "листок", "пятка", "слиток", "столик", "тяпка"}))
}

// Функция для поиска множеств анаграмм
func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	result := make(map[string][]string)

	for _, word := range words {
		lowerWord := strings.ToLower(word) // Приведение к нижнему регистру
		r := []rune(lowerWord)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		sortedWord := string(r)

		anagramMap[sortedWord] = append(anagramMap[sortedWord], lowerWord)
	}

	for _, anagrams := range anagramMap {
		if len(anagrams) > 1 {
			sort.Strings(anagrams)   // Сортировка анаграмм
			firstWord := anagrams[0] // Первое слово в множестве
			result[firstWord] = anagrams
		}
	}

	return result
}
