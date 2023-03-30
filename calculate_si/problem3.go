package main

import "fmt"

func main() {
	var Principal int
	var interestRate float32
	var noOfYears float32
	var simpleInterest float32
	fmt.Println("Enter the principal amount:")
	fmt.Scan(&Principal)
	fmt.Println("Enter the Number Of Years:")
	fmt.Scan(&noOfYears)
	fmt.Println("Enter the Rate of Interest :")
	fmt.Scan(&interestRate)
	simpleInterest = (float32(Principal) * noOfYears * interestRate) / 100
	fmt.Printf("The simple interest is %v", simpleInterest)
}
