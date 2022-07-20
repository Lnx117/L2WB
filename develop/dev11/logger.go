package main

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

//Реализуем интерфейс http.Handler и добавляем свою обработку ошибок
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handler: handlerToWrap}
}

//переписываем вывод ошибок.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//Выполнить запрос дальше, то есть передать на выполнение обработчику
	l.handler.ServeHTTP(w, r)
	//вывести лог
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}
