package main

import "errors"

func parse(req Request) (ParsedData, error) {
	if req.DriverID == "" || req.VehicleID == "" || req.LocationID == "" {
		return ParsedData{}, errors.New("driver, vehicle and location are all required")
	}
	return ParsedData{
		driverID:   req.DriverID,
		vehicleID:  req.VehicleID,
		locationID: req.LocationID,
	}, nil
}
