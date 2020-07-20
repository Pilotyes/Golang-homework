// Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в рублях и выводит сумму в долларах. Курс доллара задайте константой.
package main

import "fmt"

const usd float64 = 70.4

func main() {
	var rub float64

	fmt.Printf("Введите сумму в рублях: ")
	fmt.Scanf("%f", &rub)

	fmt.Println("Количество долларов:", rub/usd)
}
