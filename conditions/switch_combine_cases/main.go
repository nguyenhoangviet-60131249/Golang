package main

import (
	"fmt"
)

func main() {
	switch dayOfWeek := 5; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Normal days of week")
	case 6, 7:
		fmt.Println("Weekend day")
	default:
		fmt.Println("Invalid day")
	}
}
