package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	fmt.Print("Enter for weight: ")
	fmt.Scan(&weight)
	fmt.Print("Enter for height: ")
	fmt.Scan(&height)
	// condition for BMI
	var BMI = weight / (height * height)
	if BMI < 15 {
		fmt.Println("Gay")
	} else if BMI >= 15 && BMI < 22 {
		fmt.Println("Binh thuong")
	} else if BMI >= 22 && BMI < 25 {
		fmt.Println("Co nguy co beo phi")
	} else {
		fmt.Println("Beo phi")
	}
}
