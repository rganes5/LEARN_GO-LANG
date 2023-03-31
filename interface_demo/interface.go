package main

import (
	"fmt"
)

// creating  an interface
type myInterface interface {
	myArea() float64
	myVolume() float64
}

// creating a struct
type myStruct struct {
	length  float64
	breadth float64
	height  float64
}

// implementing methods of interface
func (my myStruct) myArea() float64 {
	return my.length * my.breadth
}

func (my myStruct) myVolume() float64 {
	return my.length * my.breadth * my.height
}

// main function
func main() {
	//accessing the elements of interface
	var m myInterface
	var my myStruct
	fmt.Println("Enter the length, breadth and height :")
	fmt.Scan(&my.length, &my.breadth, &my.height)
	m = my
	fmt.Printf("The area is,%v", m.myArea())
	fmt.Printf("The area is,%v", m.myVolume())
}
