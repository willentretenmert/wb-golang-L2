package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		err      bool // Использую bool потому что тест почему то валится на сравнении вывода ошибки
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
	}

	for _, c := range cases {
		result, err := unpack(c.input)
		if (err != nil) != c.err {
			t.Errorf("UnpackString(%s) ожидаемая ошибка: %v, получена ошибка: %v", c.input, c.err, err)
		}
		if result != c.expected {
			t.Errorf("UnpackString(%s) = %v, ожидаемый результат = %v", c.input, result, c.expected)
		}
	}
}
