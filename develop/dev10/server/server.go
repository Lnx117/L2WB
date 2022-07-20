package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listening: %v\n", err)
		return
	}
	defer ln.Close()

	for {
		fmt.Println("Waiting for connection...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error accepting: %v\n", err)
			return
		}
		fmt.Println("New connection. Waiting for messages...")

		go func(conn net.Conn) { // обрабатываем в го рутине чтобы мы могли общаться с кучей клиентов
			defer conn.Close()

			fmt.Printf("Serving new conn %v\n", conn)

			connReader := bufio.NewReader(conn) // ридер создается один раз

			for {
				message, err := connReader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						fmt.Printf("Connection %v closed.\n", conn)
						break
					}
					fmt.Fprintf(os.Stderr, "error reading from conn: %v\n", err)
					break
				}
				message = strings.TrimSpace(message) // удаляем \n и \r

				fmt.Printf("From: %v Received: %s\n", conn, string(message))

				newmessage := strings.ToUpper(message)
				_, err = conn.Write([]byte(newmessage + "\n"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "error writing to conn: %v\n", err)
					break
				}
			}

			fmt.Printf("Done serving client %v\n", conn)
		}(conn)
	}
}
