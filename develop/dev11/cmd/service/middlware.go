package service

import (
	"log"
	"net/http"
)

func logger(handler func(writer http.ResponseWriter, request *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		makeLogging(r)
		handler(w, r)
	}
}

func makeLogging(r *http.Request) {
	log.Println("Request: Method " + r.Method + ", Url " + r.RequestURI)
}
