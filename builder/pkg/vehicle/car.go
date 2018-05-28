package vehicle

import "fmt"

func newCar(VehicleName, paintColor, wheelType string, topVehicleSpeedMPH float64) *Car {
	myCar := new(Car)
	myCar.Properties.Name = VehicleName
	myCar.Properties.Color = paintColor
	myCar.Properties.SpeedMPH = topVehicleSpeedMPH
	myCar.Properties.SpeedKPH = topVehicleSpeedMPH * KPH
	myCar.Properties.Wheels = wheelType
	fmt.Println("Your brand new Vehicle has been built! Specs below:")
	fmt.Println(myCar)

	return myCar
}

// Drive - Drive your car
func (myCar *Car) Drive() error {

	if myCar.Properties.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myCar.Properties.SpeedMPH)
	}

	fmt.Printf("Driving your Car at max speed: %v MPH || %v KPH\n", myCar.Properties.SpeedMPH, myCar.Properties.SpeedKPH)

	return nil
}

// Stop - Stop your car
func (myCar *Car) Stop() error {

	if myCar.Properties.SpeedMPH == 0 {
		return fmt.Errorf("Can't stop Car, already stopped")
	}

	myCar.Properties.SpeedMPH = 0
	myCar.Properties.SpeedKPH = 0

	fmt.Println("Stopping car - speed is now 0")

	return nil
}

// String - print something
func (myCar *Car) String() string {

	printString := fmt.Sprintf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myCar.Properties.Name,
		myCar.Properties.Color,
		myCar.Properties.SpeedMPH,
		myCar.Properties.SpeedKPH,
		myCar.Properties.Wheels)

	return printString
}
