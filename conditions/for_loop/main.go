package main

import "fmt"

func main() {

	// print i is even number
	fmt.Println("Even Number")
	for i := 0; i <= 10; i++ {
		if i > 0 && i%2 == 0 {
			fmt.Printf("%5d", i)
		}
	}

	// print i % 3 and % 5
	fmt.Println("\nPrint number divisible 3 and 5")
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%5d", i)
		}
	}
}
