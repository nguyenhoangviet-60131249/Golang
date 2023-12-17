package main

import (
	"fmt"
	"strconv"
)

func main() {
	// declare variables
	const myFavLanguage = "Python and Golang"

	// declare multi variables
	const myName, myAge = "Nguyen Hoang Viet", 24
	const (
		badgeID string = "B221064"
	)
	fmt.Println("Name: "+myName, "\nAge: "+strconv.Itoa(myAge), "\nProgramming Language: "+myFavLanguage, "\nBadgeID: "+badgeID)

	// type constant
	const typeInt int = 20
	const typeString string = "Viet"

	fmt.Println(typeString, typeInt)
}
