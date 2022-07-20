package main

func main() {

}

//Структура компьютера которую будем возвращать
type Computer struct {
	CPU string
	RAM int
	MB  string
}

//Интерфейс описывающий любой компьютер который мы строим
//Каждый метод возвращает интерфейс чтобы можно было настраивать конфигурацию через точку
type ComputerBuilderInterface interface {
	CPU(val string) ComputerBuilderInterface
	RAM(val int) ComputerBuilderInterface
	MB(val string) ComputerBuilderInterface
	Build() Computer
}

//Структура реализующая интерфейс
type ComputerBuilder struct {
	cpu string
	ram int
	mb  string
}

func (c *ComputerBuilder) CPU(val string) ComputerBuilderInterface {
	c.cpu = val
	return c
}
func (c *ComputerBuilder) RAM(val int) ComputerBuilderInterface {
	c.ram = val
	return c
}
func (c *ComputerBuilder) MB(val string) ComputerBuilderInterface {
	c.mb = val
	return c
}

//Возвращаем финальный объект компьютера
func (c *ComputerBuilder) Build() Computer {
	return Computer{
		CPU: c.cpu,
		RAM: c.ram,
		MB:  c.mb,
	}
}

/* Преимущества и недостатки
+ Позволяет создавать продукты пошагово.
+ Позволяет использовать один и тот же код для создания различных продуктов.
+ Изолирует сложный код сборки продукта от его основной бизнес-логики.

-Усложняет код программы из-за введения дополнительных классов

Паттерн Строитель также используется, когда нужный продукт сложный и требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного
нестабильного продукта. Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.

В этом примере мы можем создать компьютеры с разными параметрами */
