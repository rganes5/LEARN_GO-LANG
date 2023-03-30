package main

import (
	"fmt"
)

func main() {
	var mainString string
	flag := 0
	fmt.Println("Enter the string")
	fmt.Scan(&mainString)
	stringLength := len(mainString)
	for i := 0; i < stringLength/2; i++ {
		if mainString[i] == mainString[stringLength-1-i] {
			flag = 0
		} else {
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Printf("The string %v is palindrome", mainString)
	} else {
		fmt.Printf("The string %v is not palindrome", mainString)
	}
}
