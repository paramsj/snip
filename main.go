package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := NewServer()
	log.Printf("snip running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, srv))
}
