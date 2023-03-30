package main

import "fmt"

func main() {
	var income uint
	fmt.Println("Enter your income: ")
	fmt.Scan(&income)
	incomeTax(income)
}
func incomeTax(income uint) {
	var tax uint
	if income <= 250000 {
		fmt.Println("There is no tax to be paid for you salary.")
	} else if income > 250000 && income <= 500000 {
		tax = (income * 5) / 100
		fmt.Printf("Income tax amount =%v", tax)
	} else if income > 500000 && income <= 1000000 {
		tax = (income * 20) / 100
		fmt.Printf("Income tax amount =%v", tax)
	} else if income > 1000000 && income <= 5000000 {
		tax = (income * 30) / 100
		fmt.Printf("Income tax amount =%v", tax)
	} else {
		fmt.Println("Contact Income Tax Department")
	}
}
