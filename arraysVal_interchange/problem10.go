package main

import (
	"fmt"
)

func main() {
	var slice1, slice2 []int
	var limit, value, temp int
	fmt.Println("Enter a limit ")
	fmt.Scan(&limit)
	fmt.Println("Enter the first array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&value)
		slice1 = append(slice1, value)
	}
	fmt.Println("Enter the second array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&value)
		slice2 = append(slice2, value)
	}
	fmt.Println("This is the swapped arrays")
	for i := 0; i < limit; i++ {
		temp = slice1[i]
		slice1[i] = slice2[i]
		slice2[i] = temp
	}
	fmt.Printf("Array 1: %v", slice1)
	fmt.Println("")
	fmt.Printf("Array 2: %v", slice2)

}
