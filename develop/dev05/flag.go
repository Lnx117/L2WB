package main

import "strings"

type Core struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Phrase     string
	CountMatch int
}

func NewCore() *Core {
	return &Core{
		After:      0,
		Before:     0,
		Context:    0,
		Count:      false,
		IgnoreCase: false,
		Invert:     false,
		Fixed:      false,
		LineNum:    false,
		Phrase:     "",
		CountMatch: 0,
	}
}

//Перевод искомой строки в нижний регистр. Нужно когда регистр неважен
func (c *Core) PhraseToLower() {
	c.Phrase = strings.ToLower(c.Phrase)
}

//Обновляем параметры after и before в зависимости от параметра context
func (c *Core) SyncOutLength() {
	if c.Context > c.After {
		c.After = c.Context
	}
	if c.Context > c.Before {
		c.Before = c.Context
	}
}

//Инкрементируем количество совпадений
func (c *Core) AddMatch() {
	c.CountMatch++
}
