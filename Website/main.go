package main

import (
	"log"
	"net/http"
)

func main() {

	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views/"))))

	log.Fatal(http.ListenAndServe(":3000", mux))
}
