// Написать функцию, которая определяет, четное ли число.
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the number ")
	fmt.Scanf("%d", &number)
	if isEven(number) == true {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}

func isEven(number int) bool {
	return number%2 == 0
}
