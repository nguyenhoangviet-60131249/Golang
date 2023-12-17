package main

import "fmt"

func main() {
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c=%d and %c = %U", myByte, myByte, myRune, myRune)
}
