package main

import "fmt"

func main() {
	fmt.Println("Pattern")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println("")
	}
}
