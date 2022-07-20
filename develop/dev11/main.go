package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

const dateLayout = "2006-01-02"

//Глобальная переменная для хранения событий
var storage Store = Store{events: make(map[int][]Event), mu: &sync.Mutex{}}

func main() {

	mux := http.NewServeMux()

	// POST routes.
	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)

	// GET routes.
	mux.HandleFunc("/events_for_day", EventsForDayHandler)
	mux.HandleFunc("/events_for_week", EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", EventsForMonthHandler)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	//Читаем порт из env
	port := ":8080"
	func() {
		temp := os.Getenv("PORT")
		if temp != "" {
			port = temp
		}
	}()

	//Реализация middleware. реализуем логгер
	wrappedMux := NewLogger(mux)

	log.Fatalln(http.ListenAndServe(":"+port, wrappedMux))

}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	//Получаем данные запроса и вносим в структуру события, если ошибка, то выводим ее
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Валидируем поля, если ошибка, то выводим ее
	if err := e.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Если все хорошо, то записываем событие в глобальную переменную
	if err := storage.Create(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Событие успешно создано", []Event{e}, http.StatusCreated)

	fmt.Println(storage.events)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	//получаем данные
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//валидируем поля
	if err := e.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//обновляем
	if err := storage.Update(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Событие было успешно обновлено", []Event{e}, http.StatusOK)

	fmt.Println(storage.events)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var deletedEvent *Event
	var err error
	if deletedEvent, err = storage.Delete(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Событие успешно удалено", []Event{*deletedEvent}, http.StatusOK)

	fmt.Println(storage.events)
}

//Получить события за конкретный день
func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	//ПОлучаем get параемтр user_id
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//ПОлучаем get параемтр date
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Ищем среди событий нужные с полученным id и датой
	var events []Event
	if events, err = storage.GetEventsForDay(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Запрос успешно выполнен!", events, http.StatusOK)
}

//Получить события за конкретную неделю
func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var events []Event
	if events, err = storage.GetEventsForWeek(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

//Получить события за конкретный год
func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var events []Event
	if events, err = storage.GetEventsForMonth(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

//Ответ с ошибкой
func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}

	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Ответ без ошибки
func resultResponse(w http.ResponseWriter, r string, e []Event, status int) {
	resultResponse := struct {
		Result string  `json:"result"`
		Events []Event `json:"events"`
	}{Result: r, Events: e}

	js, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
