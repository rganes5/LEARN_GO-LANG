package main

import "fmt"

func main() {
	defer fmt.Println("Ganesh") //works only before the ending curly braces (LIFO)
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") //LIFO
	fmt.Println("Hello")
	fmt.Println("World")

	myDefer()

}

//stack: 0,1,2,3,4
//with defer hello, world, fn call, 4,3,2,1,0, three, two , one
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
