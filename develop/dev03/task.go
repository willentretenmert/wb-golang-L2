package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

/*
Утилита sort
	Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
	на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

type sortOptions struct {
	column    int
	numeric   bool
	reverse   bool
	unique    bool
	byMonth   bool
	trimSpace bool
	check     bool
	human     bool
}

func main() {
	opts := sortOptions{}
	inputFile := flag.String("i", "", "Input file path")
	outputFile := flag.String("o", "", "Output file path")
	flag.IntVar(&opts.column, "k", 0, "column for sorting")
	flag.BoolVar(&opts.numeric, "n", false, "numeric sort")
	flag.BoolVar(&opts.reverse, "r", false, "reverse sort")
	flag.BoolVar(&opts.unique, "u", false, "unique sort")
	flag.BoolVar(&opts.byMonth, "M", false, "sort by month")
	flag.BoolVar(&opts.trimSpace, "b", false, "ignore trailing spaces")
	flag.BoolVar(&opts.check, "c", false, "check if data is sorted")
	flag.BoolVar(&opts.human, "h", false, "sort by human-readable numbers")

	flag.Parse()

	// Проверка на пустые аргументы
	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Input and output file paths are required")
		return
	}

	// Чтение строк из файла
	lines, err := readLines(*inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if opts.unique {
		lines = unique(lines)
	}

	// Сортировка строк
	sort.Strings(lines)

	if opts.reverse {
		lines = reverse(lines)
	}

	// Запись в выходной файл
	err = writeLines(lines, *outputFile)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// readLines читает строки из файла
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines записывает строки в файл
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// unique оставляет только уникальные значения
func unique(lines []string) []string {
	seen := make(map[string]struct{}) // Используйте map для отслеживания уникальных элементов
	result := []string{}              // Создайте новый слайс для уникальных элементов

	for _, v := range lines {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}       // Добавьте элемент в map
			result = append(result, v) // Добавьте элемент в результат, если он уникален
		}
	}
	return result
}

// reverse разворчаивает массив
func reverse(lines []string) []string {
	for i := 0; i < len(lines)/2; i++ {
		lines[i], lines[len(lines)-(i+1)] = lines[len(lines)-(i+1)], lines[i]
	}
	return lines
}
