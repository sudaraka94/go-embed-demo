package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static/*
var staticContent embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(StaticFS{staticContent})))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
