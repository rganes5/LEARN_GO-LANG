package main

import (
	"fmt"
)

type mainStruct struct {
	userSlices [][]int
	rows       int
	cols       int
}

func main() {
	var structDeff mainStruct
	fmt.Println("Enter the no of rows")
	fmt.Scan(&structDeff.rows)
	fmt.Println("Enter the no of rows")
	fmt.Scan(&structDeff.cols)
	getArray(&structDeff)
	displayArray(structDeff)
}
