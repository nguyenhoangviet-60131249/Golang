package main

import "fmt"

func returnMultiValuesAndError(a int, b int) (int, int, error) {
	return a, b, nil
}

func main() {
	num1 := 10
	num2 := 12
	c, d, err := returnMultiValuesAndError(num1, num2)
	if err != nil {
		fmt.Print("Wrong")
	} else {
		fmt.Print(c, d)
	}
}
