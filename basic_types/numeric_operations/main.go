package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 4, 5
	var result = (a + b) * (a + b) / 2
	fmt.Print(result)
	fmt.Print("\n")

	fmt.Print("\n")
	var r = 3.15
	var result2 = math.Pi * r * r
	fmt.Print(result2)
}
