// Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.

package main

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

}
