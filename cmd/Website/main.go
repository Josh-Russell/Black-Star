package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gopher")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
