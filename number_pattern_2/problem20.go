package main

import "fmt"

func main() {
	temp := 0
	fmt.Println("Pattern")
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			temp++
			fmt.Printf("%v ", temp)
		}
		fmt.Println("")
	}
}
