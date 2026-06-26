package main

import "log"

func scoreDriver(input ScoreRequest) ScoreResponse {

	// Fan out the three fetches; each goroutine drops its result on a channel.
	historyChannel := make(chan History, 1)
	vehicleChannel := make(chan Vehicle, 1)
	locationChannel := make(chan Location, 1)

	go func() { historyChannel <- fetchDriverHistory(input.DriverID) }()
	go func() { vehicleChannel <- fetchVehicleData(input.VehicleID) }()
	go func() { locationChannel <- fetchLocationData(input.LocationID) }()

	// Reading all three channels blocks until every fetch is done, so score
	// only runs once all the data is in.
	resp := score(<-historyChannel, <-vehicleChannel, <-locationChannel)
	return resp

}

func reportAccident(input AccidentRequest) {
	log.Printf("yo %s", input.DriverID)

}

func score(h History, v Vehicle, l Location) ScoreResponse {
	s := float64(h.Trips)*0.1 -
		float64(h.Accidents)*10 +
		float64(v.Year-2015)*2 -
		l.RiskFactor*25

	return ScoreResponse{
		Score:   s,
		Details: "computed from history, vehicle and location data",
	}
}
