package main

// Request is the raw input off the wire — it mirrors the JSON the client sends,
// which is why it carries `json` tags. Everything here is untrusted.
type ScoreRequest struct {
	DriverID   string `json:"driver"`
	VehicleID  string `json:"vehicle"`
	LocationID string `json:"location"`
}

type AccidentRequest struct {
	DriverID string `json:"driver"`
}

type History struct {
	DriverID  string `json:"driver_id"`
	Trips     int    `json:"trips"`
	Accidents int    `json:"accidents"`
}

type Vehicle struct {
	VehicleID string `json:"vehicle_id"`
	Make      string `json:"make"`
	Year      int    `json:"year"`
}

type Location struct {
	LocationID string  `json:"location_id"`
	RiskFactor float64 `json:"risk_factor"`
}

// Response is what we hand back to the caller.
type ScoreResponse struct {
	Score   float64 `json:"score"`
	Details string  `json:"details"`
}
