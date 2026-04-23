package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for:", r.URL.Path)
	fmt.Fprintf(w, "Hello, World!")
	log.Println(r.Header)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for:", r.URL.Path)
	fmt.Fprintf(w, "Current time: %s", time.Now().Format(time.RFC1123))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for:", r.URL.Path)
	fmt.Fprintf(w, "List of users: Alice, Bob, Charlie")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/time", timeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
