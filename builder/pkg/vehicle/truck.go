package vehicle

import "fmt"

func newTruck(VehicleName, paintColor, wheelType string, topVehicleSpeedMPH float64) *Truck {
	myTruck := new(Truck)
	myTruck.Properties.Name = VehicleName
	myTruck.Properties.Color = paintColor
	myTruck.Properties.SpeedMPH = topVehicleSpeedMPH
	myTruck.Properties.SpeedKPH = topVehicleSpeedMPH * KPH
	myTruck.Properties.Wheels = wheelType
	fmt.Println("Your brand new Vehicle has been built! Specs below:")
	fmt.Printf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myTruck.Properties.Name,
		myTruck.Properties.Color,
		myTruck.Properties.SpeedMPH,
		myTruck.Properties.SpeedKPH,
		myTruck.Properties.Wheels)
	return myTruck
}

// Drive - Drive your Truck
func (myTruck Truck) Drive() error {

	if myTruck.Properties.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myTruck.Properties.SpeedMPH)
	}

	fmt.Printf("Driving your Truck at max speed: %v MPH || %v KPH\n", myTruck.Properties.SpeedMPH, myTruck.Properties.SpeedKPH)

	return nil
}

// Stop - Stop your Truck
func (myTruck Truck) Stop() error {

	if myTruck.Properties.SpeedMPH == 0 {
		return fmt.Errorf("Can't stop Truck, already stopped")
	}

	myTruck.Properties.SpeedMPH = 0
	myTruck.Properties.SpeedKPH = 0

	fmt.Println("Stopping Truck - speed is now 0")

	return nil
}

// Check - Check you Truck is correct
func (myTruck Truck) Check() error {

	if myTruck.Properties.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myTruck.Properties.SpeedMPH)
	}

	if myTruck.Properties.Name == "" {
		fmt.Println("Your Truck has not been named - defaulting to name:", DefaultName)
		myTruck.Properties.Name = DefaultName
		return nil
	}

	if myTruck.Properties.Color == "" {
		fmt.Println("Your Truck has not been painted any color - defaulting to color:", DefaultColor)
		myTruck.Properties.Color = DefaultColor
		return nil
	}

	if myTruck.Properties.Wheels == "" {
		fmt.Println("Your Truck wheels have not been selected - defaulting to wheels:", DefaultWheels)
		myTruck.Properties.Wheels = DefaultWheels
		return nil
	}

	return nil
}

// String - print something
func (myTruck Truck) String() string {

	printString := fmt.Sprintf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myTruck.Properties.Name,
		myTruck.Properties.Color,
		myTruck.Properties.SpeedMPH,
		myTruck.Properties.SpeedKPH,
		myTruck.Properties.Wheels)

	return printString
}
