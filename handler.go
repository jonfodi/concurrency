package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func scoreHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeJSON[ScoreRequest](w, r)
	if !ok {
		return
	}

	start := time.Now()
	score := scoreDriver(req)

	// ~200ms (the slowest fetch) rather than ~450ms (all three added up).
	log.Printf("score = %s - handled in %s", score, time.Since(start))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(score)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("API is healthy")
	w.WriteHeader(http.StatusOK)

}

func reportAccidentHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into the wire-shaped Request...
	req, ok := decodeJSON[ScoreRequest](w, r)
	if !ok {
		return
	}

	start := time.Now()
	score := scoreDriver(req)

	// ~200ms (the slowest fetch) rather than ~450ms (all three added up).
	log.Printf("score = %s - handled in %s", score, time.Since(start))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(score)
}
