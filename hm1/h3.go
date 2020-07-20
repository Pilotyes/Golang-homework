// * Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.
package main

import "fmt"

func main() {
	var p, i float64

	fmt.Printf("Введите сумму вклада: ")
	fmt.Scanf("%f", &p)

	fmt.Printf("Введите процентную ставку: ")
	fmt.Scanf("%f", &i)

	fmt.Println("Итоговая сумма вкада через 5 лет:", p*(1+(i*5)/100))
}
