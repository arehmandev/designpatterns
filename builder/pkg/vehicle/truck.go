package vehicle

import "fmt"

// BuildTruck - creates a truck from specified params
func BuildTruck(TruckName, paintColor, wheelType string, topTruckSpeedMPH float64) *Truck {

	myTruck := new(Truck)
	myTruck.Name = TruckName
	myTruck.Color = paintColor
	myTruck.SpeedMPH = topTruckSpeedMPH
	myTruck.SpeedKPH = topTruckSpeedMPH * KPH
	myTruck.Wheels = wheelType

	fmt.Println("Your brand new Truck has been built! Specs below:")

	fmt.Printf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myTruck.Name,
		myTruck.Color,
		myTruck.SpeedMPH,
		myTruck.SpeedKPH,
		myTruck.Wheels)

	return myTruck
}

// Check - Check you Truck is correct
func (myTruck Truck) Check() error {

	if myTruck.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myTruck.SpeedMPH)
	}

	if myTruck.Name == "" {
		fmt.Println("Your Truck has not been named - defaulting to name:", DefaultName)
		myTruck.Name = DefaultName
		return nil
	}

	if myTruck.Color == "" {
		fmt.Println("Your Truck has not been painted any color - defaulting to color:", DefaultColor)
		myTruck.Color = DefaultColor
		return nil
	}

	if myTruck.Wheels == "" {
		fmt.Println("Your Truck wheels have not been selected - defaulting to wheels:", DefaultWheels)
		myTruck.Wheels = DefaultWheels
		return nil
	}

	return nil
}

// Drive - Drive your Truck
func (myTruck Truck) Drive() error {

	if myTruck.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myTruck.SpeedMPH)
	}

	fmt.Printf("Driving your Truck at max speed: %v MPH || %v KPH\n", myTruck.SpeedMPH, myTruck.SpeedKPH)

	return nil
}

// Stop - Stop your Truck
func (myTruck Truck) Stop() error {

	if myTruck.SpeedMPH == 0 {
		return fmt.Errorf("Can't stop Truck, already stopped")
	}

	myTruck.SpeedMPH = 0
	myTruck.SpeedKPH = 0

	fmt.Println("Stopping Truck - speed is now 0")

	return nil
}
