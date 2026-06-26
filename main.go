package main

import (
	"log"
	"net/http"
)

// ---------------------------------------------------------------------------
// HTTP server
// ---------------------------------------------------------------------------

func main() {
	http.HandleFunc("GET /score", scoreHandler)
	// rest of endpoints registered here
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("POST /reportAccident", reportAccidentHandler)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
