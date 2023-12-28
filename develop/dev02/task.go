package main

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
	"a4bc2d5e" => "aaaabccddddde"
	"abcd" => "abcd"
	"45" => "" (некорректная строка)
	"" => ""
Дополнительно реализовать поддержку escape-последовательностей.
Например:
	qwe\4\5 => qwe45 (*)
	qwe\45 => qwe44444 (*)
	qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(unpack("a4bc2d5e"))
	fmt.Println(unpack("abcd"))
	fmt.Println(unpack("ab3cd2"))
	fmt.Println(unpack(`qwe\4\5`))
	fmt.Println(unpack(`qwe\45`))
	fmt.Println(unpack(`qwe\\5`))
}

func unpack(s string) (string, error) {
	var result string
	var currentRune rune
	var repeatCount string
	var isEscapeSequence bool

	for _, r := range s {
		if isEscapeSequence {
			isEscapeSequence = false
			if unicode.IsDigit(r) || r == '\\' {
				currentRune = r
			} else {
				switch r {
				case 'n':
					currentRune = '\n'
				default:
					// Если после \ идет не escape-символ, возвращаем ошибку
					return "", fmt.Errorf("неверная escape последовательность \\%c", r)
				}
			}
		} else if r == '\\' {
			if currentRune != 0 {
				// Добавление предыдущего символа в результат
				count, err := strconv.Atoi(repeatCount)
				if err != nil {
					count = 1
				}
				result += strings.Repeat(string(currentRune), count)
				repeatCount = ""
			}
			currentRune = 0
			isEscapeSequence = true
			continue
		} else if unicode.IsDigit(r) {
			if currentRune == 0 && !isEscapeSequence {
				return "", fmt.Errorf("некорректная строка")
			}
			repeatCount += string(r)
		} else {
			if currentRune != 0 {
				count, err := strconv.Atoi(repeatCount)
				if err != nil {
					count = 1
				}
				result += strings.Repeat(string(currentRune), count)
				repeatCount = ""
			}
			currentRune = r
		}
	}

	if isEscapeSequence {
		return "", fmt.Errorf("незаконченная escape последовательность")
	}

	if currentRune != 0 {
		count, err := strconv.Atoi(repeatCount)
		if err != nil {
			count = 1
		}
		result += strings.Repeat(string(currentRune), count)
	}

	return result, nil
}
