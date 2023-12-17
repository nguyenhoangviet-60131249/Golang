package main

import "fmt"

func main() {

	// type interface
	var firstName, middleName, lastName, age, gender = "Nguyen", "Hoang", "Viet", 20, "Male"

	// print type of variables
	fmt.Printf("firstName :%T , middleName: %T, lastName: %T, age: %T, gender: %T", firstName, middleName, lastName, age, gender)
}
