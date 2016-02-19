package main

import (
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
