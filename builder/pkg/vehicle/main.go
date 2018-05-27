package vehicle

// Truck -
type Truck Car

// Car - struct containing all fields for a car
type Car struct {
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

// Actions - interface for car functionality
type Actions interface {
	Check() error
	Drive() error
	Stop() error
}
