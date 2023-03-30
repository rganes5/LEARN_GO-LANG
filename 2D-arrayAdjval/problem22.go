package main

import "fmt"

var userSlices1, userSlices2, sumSlices [][]int
var rows, cols int

func main() {

	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of columns")
	fmt.Scan(&cols)
	getArray()
	addArray()
	displayArray()
}

func getArray() {
	var value int
	fmt.Println("Enter the 1st array elements: ")
	for i := 0; i < rows; i++ {
		var tempSlice = []int{}
		for j := 0; j < cols; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		userSlices1 = append(userSlices1, tempSlice)
	}
	fmt.Println("Enter the 2nd array elements: ")
	for i := 0; i < rows; i++ {
		var tempSlice = []int{}
		for j := 0; j < cols; j++ {
			fmt.Scan(&value)
			tempSlice = append(tempSlice, value)
		}
		userSlices2 = append(userSlices2, tempSlice)
	}
}

func addArray() {
	var value int
	for i := 0; i < rows; i++ {
		var tempSlice = []int{}
		for j := 0; j < cols; j++ {
			value = userSlices1[i][j] + userSlices2[i][j]
			tempSlice = append(tempSlice, value)
		}
		sumSlices = append(sumSlices, tempSlice)
	}

}
func displayArray() {
	fmt.Printf("The sum array is %v", sumSlices)

}
