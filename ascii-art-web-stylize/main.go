package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "piscine/handlers"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("use go run main.go\n")
		os.Exit(1)
	}

	file := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", file))

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii", handler.AsciiArtHandler)
	http.HandleFunc("/404", handler.NotFoundHandler)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.NotFound(w, r)
	// })
	log.Println("http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
