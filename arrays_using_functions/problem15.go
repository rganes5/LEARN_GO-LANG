package main

import (
	"fmt"
)

func main() {
	userSlices := []int{}
	var limit int
	fmt.Printf("Enter the array limit")
	fmt.Scan(&limit)
	getArray(limit, &userSlices)
	displayArray(userSlices)
}

func getArray(limit int, userSlices *[]int) {
	var value int
	fmt.Printf("Enter the array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&value)
		*userSlices = append(*userSlices, value)
	}
}
func displayArray(userSlices []int) {
	fmt.Printf("The array is :%v", userSlices)

}
