package main

import (
	"fmt"
	"sort"
)

func main() {
	var limit, value int
	var userSlice []int
	fmt.Println("Enter the array limit")
	fmt.Scan(&limit)
	fmt.Println("Enter the array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&value)
		userSlice = append(userSlice, value)
	}
	fmt.Println("This is the  sorted array in descending order")
	sort.SliceStable(userSlice, func(i, j int) bool { return userSlice[i] > userSlice[j] })
	fmt.Printf("Sorted in descending order : %v", userSlice)

}
