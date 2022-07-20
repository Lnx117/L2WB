package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var fields int
	var delimiter string
	var separated bool

	flag.IntVar(&fields, "f", 0, "'fields' - выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()
	args := flag.Args()

	if fields == 0 {
		log.Fatalln("you must use -f with some value > 0")
	}

	//Считываем со строки если нет файла
	if len(args) == 0 {
		readFromTerminal(fields, delimiter, separated)
	}

	//Если файл есть, то читаем его
	fileName := args[len(args)-1]
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	//сплитим файл по строкам, а потом используем cut для каждой из них
	splitString := strings.Split(string(file), "\n")
	// проходим по всем строкам и для каждой вызываем метод Cut
	for _, str := range splitString {
		if res, ok := Cut(str, fields, delimiter, separated); ok {
			fmt.Println(res)
		}
	}
}

//Читаем пока не завершим программу
func readFromTerminal(fields int, delimiter string, separated bool) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		//Вырезаем по колонке (передаем флаги)
		res, _ := Cut(text, fields, delimiter, separated)
		fmt.Println(res)
	}
}

func Cut(str string, fields int, delimiter string, separated bool) (string, bool) {
	// ПРоверка на флаг -s, пропускаем строки без разделителя
	if separated && !strings.Contains(str, delimiter) {
		return "", false
	}
	//Сплитим строку разделителем и выводим нужный столбец
	splitStr := strings.Split(str, delimiter)
	if fields <= len(splitStr) {
		return splitStr[fields-1], true
	}
	return "", false
}
