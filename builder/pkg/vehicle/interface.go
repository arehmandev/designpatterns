package vehicle

// Actions - interface for Vehicle functionality
type Actions interface {
	Check() error
	Drive() error
	Stop() error
}

// Factory - creates a Vehicle from specified params
func Factory(VehicleName, paintColor, wheelType string, topVehicleSpeedMPH float64, vehicleType string) Actions {

	switch vehicleType {
	case "car":
		return newCar(VehicleName, paintColor, wheelType, topVehicleSpeedMPH)
	case "truck":
		return newTruck(VehicleName, paintColor, wheelType, topVehicleSpeedMPH)
	default:
		return nil
	}
}
