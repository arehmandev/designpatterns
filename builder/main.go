package main

import (
	"log"

	"github.com/arehmandev/designpatterns/builder/pkg/vehicle"
)

func main() {

	newCar := vehicle.BuildCar("Mercedes", "green", "special", float64(102.5))
	err := vehicle.RunDemo(newCar)
	check(err)

	newTruck := vehicle.BuildTruck("Mercedes", "green", "special", float64(102.5))
	err = vehicle.RunDemo(newTruck)
	check(err)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
