package main

import (
	"fmt"
)

func main() {
	var a string = "lol" // normal string
	fmt.Println(a)

	a = "Some text" // can change a string
	fmt.Println(a)

	// NOTE: The below will return the character aschii value
	fmt.Println(a[2])         // can get specific character in string (starts at index 0)
	fmt.Println(string(a[2])) // will convert ascii to string

	// a[2] = "a" // you cannot change a specific character in a string
	a = a + " more text" // string concatanation
	fmt.Println(a)
	fmt.Println([]byte(a)) // convert string to bytes
}
