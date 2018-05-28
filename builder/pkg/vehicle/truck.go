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
func (myTruck *Truck) Drive() error {

	if myTruck.Properties.SpeedMPH <= 0 {
		return fmt.Errorf("Can't drive at this speed: %v MPH", myTruck.Properties.SpeedMPH)
	}

	fmt.Printf("Driving your Truck at max speed: %v MPH || %v KPH\n", myTruck.Properties.SpeedMPH, myTruck.Properties.SpeedKPH)

	return nil
}

// Stop - Stop your Truck
func (myTruck *Truck) Stop() error {

	if myTruck.Properties.SpeedMPH == 0 {
		return fmt.Errorf("Can't stop Truck, already stopped")
	}

	myTruck.Properties.SpeedMPH = 0
	myTruck.Properties.SpeedKPH = 0

	fmt.Println("Stopping Truck - speed is now 0")

	return nil
}

// String - print something
func (myTruck *Truck) String() string {

	printString := fmt.Sprintf("[NAME: %v] Color: %v, TOPSPEED(MPH/KPH): %v/%v, Wheels: %v\n",
		myTruck.Properties.Name,
		myTruck.Properties.Color,
		myTruck.Properties.SpeedMPH,
		myTruck.Properties.SpeedKPH,
		myTruck.Properties.Wheels)

	return printString
}
