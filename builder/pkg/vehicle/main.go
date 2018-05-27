package vehicle

const (
	// DefaultName - the name of a default car
	DefaultName = "Nissan"
	// DefaultColor - the color of your car if its not selected
	DefaultColor = "silver"
	// DefaultWheels - the wheels of your car if its not selected
	DefaultWheels = "flintstonewheels"
	// MPH - Speed in Miles per hour
	MPH = 1
	// KPH - Speed of Kilometres n miles per hour
	KPH = 1.60934
)

// Car -
type Car struct {
	Properties Vehicle
}

// Truck -
type Truck struct {
	Properties Vehicle
	Height     int
	Weight     int
}

// Vehicle - struct containing all fields for a transport vehicle
type Vehicle struct {
	// Name - give a name to your car
	Name string

	// Color - color of car paint
	Color string

	// Wheels - the type of wheels to put on the car
	Wheels string

	// SpeedinMPH - top speed of car travel in MPH
	SpeedMPH float64

	// SpeedinKPH - top speed of car travel in KPH
	SpeedKPH float64
}

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
