package main

import (
	"First-Go-Project/cms"
	"net/http"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.ListenAndServe(":3010", nil)
}