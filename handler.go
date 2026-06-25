package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func scoreHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into the wire-shaped Request...
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	input, err := parse(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	start := time.Now()
	score := scoreDriver(input)

	// ~200ms (the slowest fetch) rather than ~450ms (all three added up).
	log.Printf("score = %s - handled in %s", score, time.Since(start))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(score)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("API is healthy")
	w.WriteHeader(http.StatusOK)

}
