package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("Enter the marks ")
	fmt.Scan(&mark)
	if mark >= 50 {
		fmt.Println("You are passed")
	} else {
		fmt.Println("You are failed")
	}

}
