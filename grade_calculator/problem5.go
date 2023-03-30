package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("Enter the percentage obtained: ")
	fmt.Scan(&mark)
	if mark >= 90 {
		fmt.Println("Your grade is A")
	}
	if mark >= 80 && mark <= 89 {
		fmt.Println("Your grade is B")
	}
	if mark >= 70 && mark <= 79 {
		fmt.Println("Your grade is C")
	}
	if mark >= 60 && mark <= 69 {
		fmt.Println("Your grade is D")
	}
	if mark >= 50 && mark <= 59 {
		fmt.Println("Your grade is E")
	}
	if mark < 50 {
		fmt.Println("FAILED")
	}

}
