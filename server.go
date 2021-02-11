package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, Salsa 🎉")
	}

	http.HandleFunc("/", helloHandler)

	log.Println("Running server at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
