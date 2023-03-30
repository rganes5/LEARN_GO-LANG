package main

import "fmt"

func main() {
	var number uint
	sum := 0
	fmt.Println("Enter the limit")
	fmt.Scan(&number)
	for i := 1; i <= int(number); i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Printf("The sum of odd numbers: %v", sum)

}
