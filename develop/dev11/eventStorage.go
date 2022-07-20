package main

import (
	"fmt"
	"sync"
	"time"
)

//Хранилище событий
type Store struct {
	mu     *sync.Mutex
	events map[int][]Event
}

// Создать событие, добавляет событие в хранилище
func (s *Store) Create(e *Event) error {
	//Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	if events, ok := s.events[e.UserID]; ok {
		for _, event := range events {
			if event.EventID == e.EventID {
				return fmt.Errorf("event with such id (%v) already present for this user (%v);", e.EventID, e.UserID)
			}
		}
	}

	s.events[e.UserID] = append(s.events[e.UserID], *e)

	return nil
}

//Обновляет существующее событие
func (s *Store) Update(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	//массив для хранения массива с событиями пользователя
	events := make([]Event, 0)
	ok := false

	//Проверка наличия события с нужным id пользователя
	if events, ok = s.events[e.UserID]; !ok {
		return fmt.Errorf("Пользователь с таким id (%v) не существует", e.UserID)
	}

	//Если пользоваетль есть то ищем среди событий именно этого пользователя (в events будут находиться события пользователя с
	//нужным id)
	for idx, event := range events {
		if event.EventID == e.EventID {
			index = idx
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("У пользователя с id = (%v) нет события с id = (%v)", e.UserID, e.EventID)
	}

	s.events[e.UserID][index] = *e

	return nil
}

//Удаляет событие
func (s *Store) Delete(e *Event) (*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	//массив для хранения массива с событиями пользователя
	events := make([]Event, 0)
	ok := false

	//Проверяем наличие пользователя с подобным id
	if events, ok = s.events[e.UserID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", e.UserID)
	}

	//проверяем наличие события с подобным id
	for idx, event := range events {
		if event.EventID == e.EventID {
			index = idx
			break
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("У пользователя с id = (%v) нет события с id = (%v)", e.UserID, e.EventID)
	}

	eventsLength := len(s.events[e.UserID])
	deletedEvent := s.events[e.UserID][index]
	s.events[e.UserID][index] = s.events[e.UserID][eventsLength-1]
	s.events[e.UserID] = s.events[e.UserID][:eventsLength-1]

	return &deletedEvent, nil
}

// Получить события за конкретный день
func (s *Store) GetEventsForDay(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	events := make([]Event, 0)
	ok := false

	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", userID)
	}

	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}

	return result, nil
}

//Получить события за неделю
func (s *Store) GetEventsForWeek(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	events := make([]Event, 0)
	ok := false

	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", userID)
	}

	for _, event := range events {
		y1, w1 := event.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, event)
		}
	}

	return result, nil
}

//Получить события за месяц
func (s *Store) GetEventsForMonth(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	events := make([]Event, 0)
	ok := false

	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", userID)
	}

	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}

	return result, nil
}
