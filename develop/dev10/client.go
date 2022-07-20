package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	exitChan := make(chan os.Signal, 1)
	//При нажатии ctrl + D (SIGQUIT) в канал sigCh будет отправлено сообщение
	signal.Notify(exitChan, syscall.SIGQUIT)
	go Exit(exitChan)

	//Устанавливаем флаги на задержку и иргументы на хост и порт
	timeout := flag.String("timeout", "10s", "timeout for a connection")
	flag.Parse()
	if len(flag.Args()) == 0 {
		return
	}
	//Переводим задержку в нужный формат
	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		return
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	hostPort := host + ":" + port

	// Подключаемся к сокету. Задаем таймаут подключения
	conn, err := net.DialTimeout("tcp", hostPort, timeoutDuration)
	if err != nil {
		fmt.Println("Ошибка подключения к ", hostPort)
		return
	}
	defer conn.Close()

	//Создаем ридеры
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("Ваше сообщение: ")
		text, err := console.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка считывания сообщения: %v", err)
			return
		}
		text = strings.TrimSpace(text) // удаляем \n

		//Отправляем сообщение
		fmt.Fprintf(conn, text+"\n")
		if text == "exit" {
			fmt.Println("Закрываем соединение")
			return
		}

		//Получаем ответ
		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения ответа от сервера: %v", err)
			return
		}
		// удаляем \n
		message = strings.TrimSpace(message)
		fmt.Printf("От сервера: %s\n", message)
	}
}

//Ждем нужного сигнала для выхода из программы
func Exit(exitChan chan os.Signal) {
	for {
		switch <-exitChan {
		case syscall.SIGQUIT:
			fmt.Println("User press ctrl + D .... exit")
			os.Exit(0)
		default:
		}
	}
}
