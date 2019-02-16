package main

import (
	"fmt"
)

type Vehicle interface {
	run()
	getType() string
}

type GasCar struct {

}

func (gasCar *GasCar) run() {
	fmt.Println("Burn fuel to run")
}

func (gasCar *GasCar) getType() string {
	return "Gasoline powered vehicle"
}

type BatteryCar struct {

}

func (batteryCar *BatteryCar) run() {
	fmt.Println("Drain lithium battery to run")
}

func (batteryCar *BatteryCar) getType() string {
	return "Lithium battery powered vehicle"
}

type Driver struct {
	vehicle Vehicle
}


func (driver *Driver) buyAVehicle(vehicleType int) {
	if vehicleType == 1 {
		driver.vehicle = new(GasCar)
	} else {
		driver.vehicle = new(BatteryCar)
	}
}

func (driver *Driver) driveVehicle() {
	driver.vehicle.run()
}

func (driver *Driver) showVehicle() {
	fmt.Println(driver.vehicle.getType())
}

func main() {
	var John = new(Driver)
	John.buyAVehicle(1)
	John.driveVehicle()
	John.showVehicle()

	var Terry = new(Driver)
	Terry.buyAVehicle(2)
	Terry.driveVehicle()
	Terry.showVehicle()
}
