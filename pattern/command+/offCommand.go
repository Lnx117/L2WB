package main

//Вторая команда. Получает device и реализует интерфейс command.
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}
