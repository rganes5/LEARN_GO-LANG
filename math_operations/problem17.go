package main

import (
	"fmt"
)

func main() {
	var operation, number1, number2, answer uint
	fmt.Println("Enter two numbers:")
	fmt.Scan(&number1, &number2)
	fmt.Println("Select 1 for addition")
	fmt.Println("Select 2 for Substraction")
	fmt.Println("Select 3 for multiplication")
	fmt.Println("Select 4 for division")
	fmt.Scan(&operation)
	if operation == 1 {
		addition(number1, number2, answer)
	} else if operation == 2 {
		Substraction(number1, number2, answer)
	} else if operation == 3 {
		multiplication(number1, number2, answer)
	} else if operation == 4 {
		division(number1, number2, answer)
	} else {
		fmt.Println("Invalid entry")
	}

}
func addition(number1 uint, number2 uint, answer uint) {
	answer = number1 + number2
	fmt.Printf("The answer is %v", answer)
}
func Substraction(number1 uint, number2 uint, answer uint) {
	answer = number1 - number2
	fmt.Printf("The answer is %v", answer)
}
func multiplication(number1 uint, number2 uint, answer uint) {
	answer = number1 * number2
	fmt.Printf("The answer is %v", answer)
}
func division(number1 uint, number2 uint, answer uint) {
	answer = number1 / number2
	fmt.Printf("The answer is %v", answer)
}
