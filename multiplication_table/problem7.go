package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a number")
	fmt.Scan(&number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v \n", i, number, i*number)
	}

}
