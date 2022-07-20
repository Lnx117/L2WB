package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	x := []string{"листок", "пятак", "пятка", "слиток", "столик", "тяпка"}
	y := FindAnagram(&x)
	fmt.Println(*y)
}

func FindAnagram(input *[]string) *map[string][]string {
	//массив с результатом
	result := make(map[string][]string)
	//Соответствие слова и его отсортированного массива букв
	wordAndSortedWord := make(map[string]string)
	//Пропускаем слова из одной буквы и используем функцию для поиска анаграмм
	for _, str := range *input {
		if len(str) < 2 {
			continue
		}
		anagramFinder(str, &wordAndSortedWord, &result)
	}

	//Убираем пустые подмножества (есть только одно слово, но других из таких же букв нет)
	for key, value := range result {
		if len(value) == 0 {
			delete(result, key)
		}
	}
	return &result
}

func anagramFinder(word string, wordAndSortedWord *map[string]string, result *map[string][]string) {
	//Берем слово и переводим в нижний регистр
	lower := strings.ToLower(word)

	//Слово в нижнем регистре
	wordLowerCase := []rune(lower)
	//Отсортированное посимвольно слово
	sortedWord := []rune(lower)
	//Сортируем символы в этом слове по возрастанию
	sort.Slice(sortedWord, func(i, j int) bool {
		return sortedWord[i] <= sortedWord[j]
	})
	//Теперь ищем такую комбинацию букв в номинальном массиве, если такая уже есть, значит в результате уже создано такое подмножество
	//(значит есть такой ключ)
	//В ином случае заносим в nominal соответствие сортированного слова и реального слова, а также создаем ключ и подмножество
	//в массиве с результатом
	if v, ok := (*wordAndSortedWord)[string(sortedWord)]; ok {
		(*result)[v] = append((*result)[v], string(wordLowerCase))
	} else {
		(*wordAndSortedWord)[string(sortedWord)] = string(wordLowerCase)
		(*result)[string(wordLowerCase)] = []string{}
	}
}
