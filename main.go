package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go HTTP Server!")
}

func pidginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "How e be!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/pidgin", pidginHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
