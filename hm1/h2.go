// Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу. Используйте тип данных float64 и функции из пакета math.
package main

import (
	"fmt"
	"math"
)

var a, b float64 = 3, 4

func main() {
	fmt.Println("Значение площади:", (a*b)/2)

	c := math.Sqrt(math.Pow(a, 2.0) + math.Pow(b, 2.0))
	fmt.Println("Значение периметра:", a+b+c)

	fmt.Println("Значение гипотенузы:", c)
}
