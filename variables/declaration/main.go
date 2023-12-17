package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// define variables
	var name string = "Viet "
	var myAge = time.Now().Year() - 2000
	var gender = "Male"

	// define multi variables
	var (
		yearOfBirht=2000
		firstName,lastName = "Nguyen","Viet"
	)
	fmt.Print("Name: "+ name + "\nAge: " + strconv.Itoa(myAge) + " year old" + 
				"\nGender: " + gender + "\n",
			)
	fmt.Println(firstName,lastName,yearOfBirht)
}
