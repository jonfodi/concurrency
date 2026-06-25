package main 

import (
	"math/rand"
	"time"
)

func fetchDriverHistory(driverID string) History {
	time.Sleep(150 * time.Millisecond) // pretend this is a slow DB / API call
	return History{
		DriverID:  driverID,
		Trips:     rand.Intn(500),
		Accidents: rand.Intn(5),
	}
}

func fetchVehicleData(vehicleID string) Vehicle {
	time.Sleep(200 * time.Millisecond)
	return Vehicle{
		VehicleID: vehicleID,
		Make:      "Toyota",
		Year:      2015 + rand.Intn(10),
	}
}

func fetchLocationData(locationID string) Location {
	time.Sleep(100 * time.Millisecond)
	return Location{
		LocationID: locationID,
		RiskFactor: rand.Float64(),
	}
}