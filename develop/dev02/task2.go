package main

import "fmt"

func main() {
	fmt.Println(Extract("a4bc2d5e4"))
}

func Extract(v string) string {
	if len(v) == 0 {
		return ""
	} else if len(v) == 1 {
		return v
	}

	var res []rune

	letters := []rune(v)

	if isNumber(letters[0]) {
		return "некорректная строка"
	}
	//Перебираем нашу строку
	for len(letters) > 0 {
		//На последнюю итерацию. Одно число там стоять не может тк мы удаляем два символа после добавления, но может стоять два
		if len(letters) == 1 {
			if isNumber(letters[0]) {
				return "некорректная строка"
			}
			res = append(res, letters[0])
			break
		}
		//Смотрим чем являются два первых символа. Если оба числа то выводим сообщение, если буква и число то обрабатываем
		a, b := letters[0], letters[1]
		if isNumber(a) && isNumber(b) {
			return "некорректная строка"
		}
		//Если второй символ число то добавляем такое количество букв и обрезаем исходный массив
		if isNumber(b) {
			// int(b-'0') преобразует rune в int
			res = append(res, repeatedLetter(int(b-'0'), a)...)
			letters = letters[2:]
			//Если не число то просто добавляем новую букву в массив
		} else {
			res = append(res, a)
			letters = letters[1:]
		}
	}
	return string(res)

}

func isNumber(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func repeatedLetter(n int, sym rune) []rune {
	var res []rune

	for i := 0; i < n; i++ {
		res = append(res, sym)
	}

	return res
}
