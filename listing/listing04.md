package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}

/*Вывод 0-9 потом дедлок
Дедлок потому что мы продолжаем читать канал в который никто не пишет. Канал не закрыли*/

