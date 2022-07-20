package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	core := NewCore()
	flag.IntVar(&core.After, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&core.Before, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&core.Context, "C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&core.Count, "c", false, "'count' (количество строк)")
	flag.BoolVar(&core.IgnoreCase, "i", false, "'ignore-case' (игнорировать регистр)")
	flag.BoolVar(&core.Invert, "v", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&core.Fixed, "F", false, "'fixed', точное совпадение со строкой")
	flag.BoolVar(&core.LineNum, "n", false, "'line num', печатать номер строки")
	flag.Parse()
	//Обновляем параметры after и before в зависимости от параметра context
	core.SyncOutLength()
	//Собираем аргументы, искомая строка и название файла
	args := flag.Args()

	//Если аргументов нехватает пишем как использовать программу
	if len(args) < 2 {
		log.Fatalln("Чтобы начать поиск: [флаги] [искомая строка] [название файла]")
	}

	//Искомая фраза
	slicePhrase := args[:len(args)-1]
	//Объединяем в одну строку если фраза состоит из нескольких слов
	core.Phrase = strings.Join(slicePhrase, " ")

	//Считываем файл
	file, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	//Сплитим файл построчно
	splitString := strings.Split(string(file), "\n")
	//Используем grep и записываем результат, затем печатаем его
	result := Grep(splitString, core)
	printRes(core, result)
}

// Grep функция поиска фразы или строки в файле с применением доп.условий
func Grep(text []string, c *Core) []*GrepStruct {
	//Слайс с результатами поиска. Это массив из узлов, каждый из которых массив ключ значение
	var result []*GrepStruct
	var condition bool // условие сравнения

	//Проходим построчно по файлу
	for index, str := range text {
		// если применен -i, убираем регистр
		if c.IgnoreCase {
			//Переводим текущую строку в нижний регистр
			str = strings.ToLower(str)
			//Переводим искомую фразу в нижний регистр
			c.PhraseToLower()
		}
		//Проверяем условия
		if c.Fixed {
			condition = c.Phrase == str // полное совпадение строки
		} else {
			condition = strings.Contains(str, c.Phrase) // совпадение подстроки
		}

		//Флаг исключения
		if c.Invert {
			condition = !condition
		}

		//Создаем объект grep для добавления в результат
		match := NewGrep()
		// если условие выполняется то значит в эту строку записываем
		if condition {
			c.AddMatch()
			//Определяем количество строк для печати в зависимости от флагов
			var upRange, downRange = 0, len(text) - 1
			if d := index - c.Before; d > upRange {
				upRange = d
			}
			if d := index + c.After; d < downRange {
				downRange = d
			}
			for i := upRange; i <= downRange; i++ {
				match.Result = append(match.Result, Node{
					Key:   i + 1,
					Value: text[i],
				})
			}
			result = append(result, match)
		}

	}
	return result
}

//вывод результата
func printRes(c *Core, res []*GrepStruct) {
	//Если установлен флаг на количество вхождений
	if c.Count {
		fmt.Printf("Совпадений: %d\n", c.CountMatch)
	}

	//Проходим по результату
	for _, match := range res {
		match.Print(c.LineNum)
	}
}
