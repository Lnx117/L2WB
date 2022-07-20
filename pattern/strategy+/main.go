/* https://www.youtube.com/watch?v=AeynEfQ8rHg&ab_channel=TheArtofDevelopment */

package main

/* Интернет магазин с заказом и несколькими способами его оплаты */
//Payment это лишь малая часть бизнес логики
func main() {
	product := "auto"

	//Способ оплаты
	payWay := 3
	var payment Payment

	//Получаем нужный объект в зависимости от способа оплаты
	switch payWay {
	case 1:
		payment = NewCardPayment()
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}

//Функция оплаты
func processOrder(product string, payment Payment) {
	err := payment.Pay()

	if err != nil {
		return
	}
}

//Интерфейс для разных способов оплаты
type Payment interface {
	Pay() error
}

//Структуры разных способов оплаты и конструкторы
type CardPayment struct {
}

func (p *CardPayment) Pay() error {
	return nil
}

func NewCardPayment() *CardPayment {
	return &CardPayment{}
}

type PayPalPayment struct {
}

func (p *PayPalPayment) Pay() error {
	return nil
}

func NewPayPalPayment() *PayPalPayment {
	return &PayPalPayment{}
}

type QIWIPayment struct {
}

func (p *QIWIPayment) Pay() error {
	return nil
}

func NewQIWIPayment() *QIWIPayment {
	return &QIWIPayment{}
}

/* Используем чтобы в большом слое бизнес логики выделить и отделить ее часть имея малую связаннойсть (low coupling) */
