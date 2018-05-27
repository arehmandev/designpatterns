package vehicle

import "fmt"

func newCar(VehicleName, paintColor, wheelType string, topVehicleSpeedMPH float64) *Car {
	myVehicle := new(Car)
	myVehicle.Name = VehicleName
	myVehicle.Color = paintColor
	myVehicle.SpeedMPH = topVehicleSpeedMPH
	myVehicle.SpeedKPH = topVehicleSpeedMPH * KPH
	myVehicle.Wheels = wheelType
	fmt.Println("Your brand new Vehicle has been built! Specs below:")
	fmt.Println(myVehicle)

	return myVehicle
}

// Drive - Drive your car
func (myCar Car) Drive() error {

	if myCar.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myCar.SpeedMPH)
	}

	fmt.Printf("Driving your Car at max speed: %v MPH || %v KPH\n", myCar.SpeedMPH, myCar.SpeedKPH)

	return nil
}

// Stop - Stop your car
func (myCar Car) Stop() error {

	if myCar.SpeedMPH == 0 {
		return fmt.Errorf("Can't stop Car, already stopped")
	}

	myCar.SpeedMPH = 0
	myCar.SpeedKPH = 0

	fmt.Println("Stopping car - speed is now 0")

	return nil
}

// Check - Check you Car is correct
func (myCar Car) Check() error {

	if myCar.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myCar.SpeedMPH)
	}

	if myCar.Name == "" {
		fmt.Println("Your Car has not been named - defaulting to name:", DefaultName)
		myCar.Name = DefaultName
		return nil
	}

	if myCar.Color == "" {
		fmt.Println("Your Car has not been painted any color - defaulting to color:", DefaultColor)
		myCar.Color = DefaultColor
		return nil
	}

	if myCar.Wheels == "" {
		fmt.Println("Your Car wheels have not been selected - defaulting to wheels:", DefaultWheels)
		myCar.Wheels = DefaultWheels
		return nil
	}

	return nil
}

// String - print something
func (myCar Car) String() string {

	printString := fmt.Sprintf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myCar.Name,
		myCar.Color,
		myCar.SpeedMPH,
		myCar.SpeedKPH,
		myCar.Wheels)

	return printString
}
