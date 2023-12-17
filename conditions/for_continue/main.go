package main

import "fmt"

func main() {
	for num := 0; num <= 10; num++ {
		if num%2 == 0 {
			continue // if i num is even number then pass and print odd number
		}
		fmt.Printf("%5d", num)
	}
}
