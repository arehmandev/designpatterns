package vehicle

import "fmt"

func newTruck(VehicleName, paintColor, wheelType string, topVehicleSpeedMPH float64) *Truck {
	myVehicle := new(Truck)
	myVehicle.Name = VehicleName
	myVehicle.Color = paintColor
	myVehicle.SpeedMPH = topVehicleSpeedMPH
	myVehicle.SpeedKPH = topVehicleSpeedMPH * KPH
	myVehicle.Wheels = wheelType
	fmt.Println("Your brand new Vehicle has been built! Specs below:")
	fmt.Printf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myVehicle.Name,
		myVehicle.Color,
		myVehicle.SpeedMPH,
		myVehicle.SpeedKPH,
		myVehicle.Wheels)
	return myVehicle
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

// String - print something
func (myTruck Truck) String() string {

	printString := fmt.Sprintf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myTruck.Name,
		myTruck.Color,
		myTruck.SpeedMPH,
		myTruck.SpeedKPH,
		myTruck.Wheels)

	return printString
}
