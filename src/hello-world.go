package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	Port := ":8080"
	message := os.Getenv("MESSAGE")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})

	fmt.Printf("Starting server on port %s.\n", Port)
	http.ListenAndServe(Port, nil)
}
