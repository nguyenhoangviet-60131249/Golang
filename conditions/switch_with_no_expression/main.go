package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	// user input weight and height
	fmt.Print("Enter your weight: ")
	fmt.Scan(&weight)
	fmt.Printf("Enter your height: ")
	fmt.Scan(&height)
	var BMI = weight / (height * height)

	switch {
	case BMI <= 15:
		fmt.Println("Gay")
	case BMI >= 15 && BMI < 22:
		fmt.Println("Binh thuong")
	case BMI >= 22:
		fmt.Println("Beo phi")
	default:
		fmt.Println("Invalid BMI")
	}
}
