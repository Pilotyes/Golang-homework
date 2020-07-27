// Инициализировать несколько экземпляров структур. Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.

package main

import "fmt"

type car struct {
	carBrand          string
	yearOfManufacture int16
	trunk             float32
	isIgnitionEnabled bool
	isWindowsOpen     bool
	trunkPercent      float32
}
type truck struct {
	carBrand          string
	yearOfManufacture int16
	truckBody         float32
	isIgnitionEnabled bool
	isWindowsOpen     bool
	truckBodyPercent  float32
}

func main() {
	car1 := car{
		carBrand:          "Mercedes-Benz",
		yearOfManufacture: 2020,
		trunk:             600.5,
		isIgnitionEnabled: true,
		isWindowsOpen:     false,
		trunkPercent:      25,
	}
	var truck1 = truck{
		carBrand:          "Gazel",
		yearOfManufacture: 1990,
		truckBody:         3000,
		isIgnitionEnabled: false,
		isWindowsOpen:     true,
		truckBodyPercent:  1.9,
	}
	fmt.Println(car1)
	fmt.Println(truck1)

	car1.isIgnitionEnabled = false
	truck1.isIgnitionEnabled = true
	truck1.truckBodyPercent = 50

	fmt.Println()
	fmt.Println(car1)
	fmt.Println(truck1)

	car2 := car{
		carBrand:          "Nissan",
		yearOfManufacture: 2015,
		trunk:             300.5,
		isIgnitionEnabled: false,
		isWindowsOpen:     false,
		trunkPercent:      25,
	}
	fmt.Println()
	fmt.Println(car1 != car2)
}
