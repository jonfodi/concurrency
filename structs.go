package main

// Request is the raw input off the wire — it mirrors the JSON the client sends,
// which is why it carries `json` tags. Everything here is untrusted.
type Request struct {
	DriverID   string `json:"driver"`
	VehicleID  string `json:"vehicle"`
	LocationID string `json:"location"`
}

// ParsedData is the validated, internal form the rest of the code works with.
// It's a distinct type from Request on purpose: Request is "whatever the client
// sent", ParsedData is "input we've checked and trust".
type ParsedData struct {
	driverID   string
	vehicleID  string
	locationID string
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
type Response struct {
	Score   float64 `json:"score"`
	Details string  `json:"details"`
}