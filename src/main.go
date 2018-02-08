package main

import (
	"net/http"
	"shortener"
)

func init() {
	http.HandleFunc("/", shortener.Index)
	http.HandleFunc("/shortener", shortener.CreateShortUrl)
}