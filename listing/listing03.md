package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	//сравнение
	fmt.Println(err == nil)
}

/*Вывод
< nil >
false

Тип interface под капотом выглядит вот так:

type iface struct {
	//это указатель на Interface Table или itable — структуру, которая хранит некоторые метаданные о типе и список методов,
	//используемых для удовлетворения интерфейса.
    tab  *itab
	//data — указывает на фактическую переменную с конкретным (статическим) типом,
    data unsafe.Pointer
}

Поскольку у пустого интерфейса нет никаких методов, то и itable для него просчитывать и хранить не нужно — достаточно только метаинформации о статическом типе.

type iface struct {
	//data — указывает на фактическую переменную с конкретным (статическим) типом,
    data unsafe.Pointer
}


type itab struct { // 40 bytes on a 64bit arch
    inter *interfacetype
    _type *_type
    ...
}
Интерфейс хранит в себе тип интерфейса и тип самого значение.

Значение любого интерфейса, не только error, является nil в случае когда И значение И тип являются nil.

Функция Foo возвращает nil типа *os.PathError, результат мы сравниваем с nil типа nil, откуда и следует их неравенство.
*/
