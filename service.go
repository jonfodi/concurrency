package main

func scoreDriver(input ParsedData) Response {

	// Fan out the three fetches; each goroutine drops its result on a channel.
	historyChannel := make(chan History, 1)
	vehicleChannel := make(chan Vehicle, 1)
	locationChannel := make(chan Location, 1)

	go func() { historyChannel <- fetchDriverHistory(input.driverID) }()
	go func() { vehicleChannel <- fetchVehicleData(input.vehicleID) }()
	go func() { locationChannel <- fetchLocationData(input.locationID) }()

	// Reading all three channels blocks until every fetch is done, so score
	// only runs once all the data is in.
	resp := score(<-historyChannel, <-vehicleChannel, <-locationChannel)
	return resp

}

func score(h History, v Vehicle, l Location) Response {
	s := float64(h.Trips)*0.1 -
		float64(h.Accidents)*10 +
		float64(v.Year-2015)*2 -
		l.RiskFactor*25

	return Response{
		Score:   s,
		Details: "computed from history, vehicle and location data",
	}
}
