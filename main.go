package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Print("Sleeping for a few seconds before starting server...")
	time.Sleep(30 * time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello there")
	})

	log.Print("Start listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Something wrong happened: %v", err)
	}
}
