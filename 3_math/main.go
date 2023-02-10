package main

import "fmt"

func main() {
	var a int = 5
	var b int = 7

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(b / a)
	fmt.Println(a % b) // note: can only be used on integers

	var c int8 = 19

	// a + c is not allowed in Go, since int and int8 are considered
	// different, so you have to convert one to be the other
	// in this case we converted 'c' to int
	fmt.Println(a + int(c))
}
