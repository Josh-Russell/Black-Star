package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.dir("templates/"))))

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
