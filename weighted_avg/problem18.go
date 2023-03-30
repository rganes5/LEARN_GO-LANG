package main

import "fmt"

func main() {
	var wMark, lMark, aMark uint
	fmt.Println("Enter the written test mark :")
	fmt.Scan(&wMark)
	fmt.Println("Enter the Lab exam mark :")
	fmt.Scan(&lMark)
	fmt.Println("Enter the assignment mark :")
	fmt.Scan(&aMark)
	grade(wMark, lMark, aMark)
}
func grade(wMark uint, lMark uint, aMark uint) {
	var grade float32
	grade = (float32(wMark)*70)/100 + (float32(lMark)*20)/100 + (float32(aMark)*10)/100
	fmt.Printf("Grade of the student is %v", grade)
}
