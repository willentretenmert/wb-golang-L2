package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			input: []string{"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Столик"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			input:    []string{"машина", "компьютер", "телефон"},
			expected: map[string][]string{},
		},
	}

	for _, c := range cases {
		result := findAnagrams(c.input)

		// Для сравнения мап, необходимо сначала отсортировать значения
		for _, anagrams := range result {
			sort.Strings(anagrams)
		}

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Для '%v' ожидалось '%v', получено '%v'", c.input, c.expected, result)
		}
	}
}
