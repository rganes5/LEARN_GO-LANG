package main

import "fmt"

func main() {
	var value, rows, cols int
	var userSlices1, userSlices2, sumSlices [][]int
	fmt.Println("Enter the no of rows")
	fmt.Scan(&rows)
	fmt.Println("Enter the no of columns")
	fmt.Scan(&cols)
	//Input to 1st array
	fmt.Println("Enter the first 2D array elements")
	for i := 0; i < rows; i++ {
		tempSlice := []int{}
		for j := 0; j < cols; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		userSlices1 = append(userSlices1, tempSlice)
	}
	//Input to 2nd array
	fmt.Println("Enter the first 2D array elements")
	for i := 0; i < rows; i++ {
		tempSlice := []int{}
		for j := 0; j < cols; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		userSlices2 = append(userSlices2, tempSlice)
	}

	fmt.Println("The sum of the two 2D arrays are")
	for i := 0; i < rows; i++ {
		tempSlice := []int{}
		for j := 0; j < cols; j++ {
			value = userSlices1[i][j] + userSlices2[i][j]
			tempSlice = append(tempSlice, value)
		}
		sumSlices = append(sumSlices, tempSlice)
	}
	fmt.Printf("%v", sumSlices)
}
