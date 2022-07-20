package main

import "testing"

func TestExtract(t *testing.T) {
	tests := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "некорректная строка",
		"":         "",
	}

	for key, value := range tests {
		res := Extract(key)

		if res != value {
			t.Errorf("Ошибка некорректная строка")
			continue
		}
	}
}
