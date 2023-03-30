package main

import (
	"fmt"
	"math/big"
)

func main() {
	var number uint
	fmt.Println("Please enter a number: ")
	fmt.Scan(&number)
	// if primes.IsPrime(number) {
	if big.NewInt(int64(number)).ProbablyPrime(0) {
		fmt.Printf("The number %v is prime number", number)
	} else {
		fmt.Printf("The number %v is not Prime number", number)
	}
}
