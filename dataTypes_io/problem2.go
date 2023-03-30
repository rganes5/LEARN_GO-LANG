package main

import "fmt"

func main() {
	var number1 int
	var number2 float32
	var sum float32
	fmt.Println("Enter a integer: ")
	fmt.Scan(&number1)
	fmt.Println("Enter a decimal number: ")
	fmt.Scan(&number2)
	sum = float32(number1) + number2
	fmt.Printf("The sum of the numbers %v and %f is = %f", number1, number2, sum)

}
