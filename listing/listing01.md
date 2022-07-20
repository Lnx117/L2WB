package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

//Вывод [77 78 79]
//Создаем срез из массива и выводим c 1 го по 4-1 ый, т.е. 3 ий
