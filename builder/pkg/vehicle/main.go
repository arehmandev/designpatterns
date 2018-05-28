package vehicle

import "fmt"

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
	Drive() error
	Stop() error
}

// Factory - creates a Vehicle from specified params
func Factory(VehicleName, paintColor, wheelType string, topVehicleSpeedMPH float64, vehicleType string) (actions Actions, err error) {

	switch vehicleType {
	case "car":
		car := newCar(VehicleName, paintColor, wheelType, topVehicleSpeedMPH)
		err := car.Properties.Check()
		return car, err
	case "truck":
		truck := newTruck(VehicleName, paintColor, wheelType, topVehicleSpeedMPH)
		err := truck.Properties.Check()
		return truck, err
	default:
		return nil, nil
	}
}

// Check - Check you Car is correct
func (myVehicle *Vehicle) Check() error {

	if myVehicle.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myVehicle.SpeedMPH)
	}

	if myVehicle.Name == "" {
		fmt.Println("Your Car has not been named - defaulting to name:", DefaultName)
		myVehicle.Name = DefaultName
		return nil
	}

	if myVehicle.Color == "" {
		fmt.Println("Your Car has not been painted any color - defaulting to color:", DefaultColor)
		myVehicle.Color = DefaultColor
		return nil
	}

	if myVehicle.Wheels == "" {
		fmt.Println("Your Car wheels have not been selected - defaulting to wheels:", DefaultWheels)
		myVehicle.Wheels = DefaultWheels
		return nil
	}

	return nil
}
