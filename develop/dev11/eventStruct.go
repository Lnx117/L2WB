package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Структура событий
type Event struct {
	UserID      int       `json:"user_id"`
	EventID     int       `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

//Записываем пришедшие данные в структуру (переводим из json в структурку)
func (e *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//Проверяем все ли поля правильно заполнены. Нужны user_id, event_id и title
func (e *Event) Validate() error {
	if e.UserID <= 0 {
		return fmt.Errorf("Неверный user_id: %v;", e.UserID)
	}

	if e.EventID <= 0 {
		return fmt.Errorf("iНеверный event_id: %v;", e.EventID)
	}

	if e.Title == "" {
		return fmt.Errorf("title cannot be empty;")
	}

	return nil
}
