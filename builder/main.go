package main

import (
	"fmt"
	"log"

	"github.com/arehmandev/designpatterns/builder/pkg/vehicle"
)

func main() {
	fmt.Println("------- VEHICLE 1 -------")
	newCar, err := vehicle.Factory("Mercedes", "green", "special", float64(102.5), "car")
	check(err)
	err = vehicle.RunDemo(newCar)
	check(err)

	fmt.Println("------- VEHICLE 2 -------")

	newTruck, err := vehicle.Factory("BigTruck", "grey", "huge", float64(82.5), "truck")
	check(err)
	err = vehicle.RunDemo(newTruck)
	check(err)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
