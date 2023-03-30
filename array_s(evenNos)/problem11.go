package main

import "fmt"

func main() {
	var userSlice []int
	var limit, value int
	evenCount := 0
	fmt.Println("Enter the limit of the array")
	fmt.Scan(&limit)
	fmt.Println("Enter the array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&value)
		userSlice = append(userSlice, value)
	}
	for _, v := range userSlice {
		if v%2 == 0 {
			evenCount++
		}
	}
	fmt.Printf("The number of even numbers in the array is %v", evenCount)
}
