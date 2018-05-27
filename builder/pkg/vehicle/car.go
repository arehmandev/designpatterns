package vehicle

import "fmt"

// BuildCar - creates a car from specified params
func BuildCar(carName, paintColor, wheelType string, topCarSpeedMPH float64) *Car {

	myCar := new(Car)
	myCar.Name = carName
	myCar.Color = paintColor
	myCar.SpeedMPH = topCarSpeedMPH
	myCar.SpeedKPH = topCarSpeedMPH * KPH
	myCar.Wheels = wheelType

	fmt.Println("Your brand new car has been built! Specs below:")

	fmt.Printf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myCar.Name,
		myCar.Color,
		myCar.SpeedMPH,
		myCar.SpeedKPH,
		myCar.Wheels)

	return myCar
}

// Check - Check you car is correct
func (myCar Car) Check() error {

	if myCar.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myCar.SpeedMPH)
	}

	if myCar.Name == "" {
		fmt.Println("Your car has not been named - defaulting to name:", DefaultName)
		myCar.Name = DefaultName
		return nil
	}

	if myCar.Color == "" {
		fmt.Println("Your car has not been painted any color - defaulting to color:", DefaultColor)
		myCar.Color = DefaultColor
		return nil
	}

	if myCar.Wheels == "" {
		fmt.Println("Your car wheels have not been selected - defaulting to wheels:", DefaultWheels)
		myCar.Wheels = DefaultWheels
		return nil
	}

	return nil
}

// Drive - Drive your car
func (myCar Car) Drive() error {

	if myCar.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myCar.SpeedMPH)
	}

	fmt.Printf("Driving your car at max speed: %v MPH || %v KPH\n", myCar.SpeedMPH, myCar.SpeedKPH)

	return nil
}

// Stop - Stop your car
func (myCar Car) Stop() error {

	if myCar.SpeedMPH == 0 {
		return fmt.Errorf("Can't stop car, already stopped")
	}

	myCar.SpeedMPH = 0
	myCar.SpeedKPH = 0

	fmt.Println("Stopping car - speed is now 0")

	return nil
}
