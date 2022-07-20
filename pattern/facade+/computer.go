package main

import "fmt"

type Computer struct {
}

func NewComputer() *Computer {
	return &Computer{}
}

func (c *Computer) GetElectric() {
	fmt.Println("Включаю питание!")
}

func (c *Computer) GetLoadingScreen() {
	fmt.Println("Загружаю систему...")
}

func (c *Computer) SayHelloUser() {
	fmt.Println("Привет User")
}

func (c *Computer) ReadyToWork() {
	fmt.Println("Система готова к работе!")
}

func (c *Computer) StoppingAllPrograms() {
	fmt.Println("Завершаю работу всех программ")
}

func (c *Computer) getTurnOffScreen() {
	fmt.Println("Выключение компьютера...")
}

func (c *Computer) StopGettingElectric() {
	fmt.Println("Выключаю питание")
}

func X() {
	fmt.Println("Выключаю питание")
}
