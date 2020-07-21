// Написать функцию, которая определяет, делится ли число без остатка на 3.
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the number ")
	fmt.Scanf("%d", &number)
	if isDivid(number) == true {
		fmt.Println("The number is divided into three")
	} else {
		fmt.Println("The number is not divided into three")
	}
}

func isDivid(number int) bool {
	return number%3 == 0
}
