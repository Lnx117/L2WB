package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Функция создает канал, который закрывается спустя заданное время
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	//or получает список каналов и возвращает один, который завершится при завершении любого канала из списка
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))
}

func or(ch ...<-chan interface{}) <-chan interface{} {
	// создаем single канал который вернем
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)
	// Запускаем горутины прослушивающие каналы из списка(создаем каналы)
	for _, channel := range ch {
		//анонимная функция, смотрит за каналом, когда он завершится, сработает wg.Done() и мы сможем закрыть single канал
		go func(channel <-chan interface{}) {
			// в каждом из них запускаем цикл который завершится по закрытии канала
			for range channel {
			}
			wg.Done()
		}(channel)
	}

	wg.Wait()
	// закрываем сингл канал и возвращаем
	close(out)
	return out
}
