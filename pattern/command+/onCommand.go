package main

//Первая команда
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}
