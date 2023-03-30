package main

import "fmt"

func main() {
	number1 := 15
	number2 := 25
	//Swapping
	number1 = number1 + number2
	number2 = number1 - number2
	number1 = number1 - number2

	fmt.Println("The numbers after swapping :")
	fmt.Printf("number1=%v number2=%v", number1, number2)
}
