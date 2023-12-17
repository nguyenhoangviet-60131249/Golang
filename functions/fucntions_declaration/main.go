package main

import "fmt"

func avg(x int, y int) int {
	return (x + y) / 2
}
func main() {
	x := 6
	y := 10
	result := avg(x, y)
	fmt.Printf("Average of %d and %d = %d\n", x, y, result)
}
