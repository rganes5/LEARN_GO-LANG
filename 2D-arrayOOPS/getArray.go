package main

import "fmt"

func getArray(structDeff *mainStruct) {
	var v int
	fmt.Println("Enter the array elements")
	for i := 0; i < structDeff.rows; i++ {
		var tempSlice []int
		for i := 0; i < structDeff.cols; i++ {
			fmt.Scan(&v)
			tempSlice = append(tempSlice, v)
		}
		structDeff.userSlices = append(structDeff.userSlices, tempSlice)
	}
}
