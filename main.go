package main

import (
	"log"
	"net/http"
)

// ---------------------------------------------------------------------------
// HTTP server
// ---------------------------------------------------------------------------

func main() {
	http.HandleFunc("/score", scoreHandler)
	// rest of endpoints registered here
	http.HandleFunc("/health", healthCheckHandler)
	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
