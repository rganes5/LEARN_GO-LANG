package main

import "fmt"

func main() {
	var inputSlices, outputSlices []int
	var limit, value int
	fmt.Println("Enter the array limit")
	fmt.Scan(&limit)
	fmt.Println("Enter the array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&value)
		inputSlices = append(inputSlices, value)
	}
	// multiplication of adjacent elements

	for i := 0; i < limit; i++ {
		if i < limit-1 {
			outputSlices = append(outputSlices, inputSlices[i]*inputSlices[i+1])
		}
	}
	fmt.Printf("Here is the output: %v", outputSlices)
}
