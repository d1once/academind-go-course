package main

import "fmt"

func main() {
	// var firstName string = "Dionis"
	// var lastName string = "Senja"
	firstName := "Dionis"
	lastName := "Senja"
	// var currentYear int = 2021
	// var birthYear int = 1999
	// var differenceOfYears = currentYear - birthYear
	currentYear := 2021
	birthYear := 1999
	differenceOfYears := currentYear - birthYear
	currentYear += 1

	fmt.Println(firstName, lastName)
	fmt.Println(differenceOfYears)
	fmt.Println(currentYear)
}