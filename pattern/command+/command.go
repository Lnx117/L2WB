package main

//интерфейс команды
type command interface {
	execute()
}
