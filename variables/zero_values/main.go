package main

import "fmt"

func main() {

	// define multi variable same type
	var (
		firstName, middleName, lastName string
		age                             int
		gender                          string
	)

	// print type of variables
	fmt.Printf("firstName :%T , middleName: %T, lastName: %T, age: %T, gender: %T", firstName, middleName, lastName, age, gender)
}
