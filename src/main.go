package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"shortener"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", shortener.Index).Methods("GET")
	rtr.HandleFunc("/shortener", shortener.CreateShortUrl).Methods("POST")
	rtr.HandleFunc("/{urlHash:[a-zA-Z0-9]*}", shortener.OriginalRedirect)

	http.Handle("/", rtr)
}