package main

//Отправитель
//Получает объект команды
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
