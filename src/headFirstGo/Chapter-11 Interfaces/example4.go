package main

import "fmt"

type Car string
type Truck string

func (c Car) Accelerate() {
	fmt.Println("Car Speeding Up!!")
}

func (c Car) Brake() {
	fmt.Println("Car Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Car Turining", direction)
}

func (t Truck) Accelerate() {
	fmt.Println("Truck Speeding Up!!!")
}

func (t Truck) Brake() {
	fmt.Println("Truck Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Truck Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Truck Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Honda")
	vehicle.Accelerate()
	vehicle.Steer("Right - Left")

	vehicle = Truck("Optimus prime")
	vehicle.Accelerate()
	vehicle.Brake()
	Truck("Megatron").LoadCargo("transform")
}
