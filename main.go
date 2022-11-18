package main

import (
	"fmt"
	"rhin/lib"
	"rhin/pipe"
)

type Human struct {
	name   string
	gender string
}

type Car struct {
	name     string
	maxSpeed int
}

type Speed struct {
	fast bool
}

func main() {
	humans := []Human{
		{name: "Bruce", gender: "male"},
		{name: "Mia", gender: "female"},
		{name: "Vincent", gender: "male"},
		{name: "Lucy", gender: "female"},
	}

	cars := []Car{
		{name: "Mustang", maxSpeed: 200},
		{name: "Kangoo", maxSpeed: 160},
		{name: "Peugeot 205", maxSpeed: 140},
	}

	isMale := func(item Human) bool {
		return item.gender == "male"
	}
	fmt.Println("Every humans are male:", lib.Every(humans, isMale))
	fmt.Println("Some human are male:", lib.Some(humans, isMale))
	fmt.Println("Filter on males:", lib.Filter(humans, isMale))

	isSportCar := func(item Car) bool {
		return item.maxSpeed > 180
	}
	fmt.Println("Every cars are sport car:", lib.Every(cars, isSportCar))
	fmt.Println("Some car are sport car:", lib.Some(cars, isSportCar))
	fmt.Println("Filter on sport cars:", lib.Filter(cars, isSportCar))
	fmt.Println("Map cars to speeds:", lib.Map(cars, func(item Car) Speed { return Speed{fast: isSportCar(item)} }))

	res := pipe.P3(
		cars,
		pipe.Filter(func (car Car) bool {
			return car.maxSpeed > 150
		}),
		pipe.Map(func (car Car) int {
			return car.maxSpeed
		}),
		pipe.Some(func (speed int) bool {
			return speed > 100
		}),
	)

	fmt.Println(res)
}
