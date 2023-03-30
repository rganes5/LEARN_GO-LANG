package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1, string2 string
	fmt.Println("Enter the first string")
	fmt.Scan(&string1)
	fmt.Println("Enter the Second string")
	fmt.Scan(&string2)
	true := strings.Compare(string1, string2)
	if true == 0 {
		fmt.Println("They are both same")
	} else {
		fmt.Println("They are not same.")
	}

}
