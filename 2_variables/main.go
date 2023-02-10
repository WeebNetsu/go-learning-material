package main

import "fmt"

// when in global scope, you cannot use := syntax
// note: with scoping in go, the below is global, but only to this file
var c int = 100
var H int = 1 // this is global scope to files outside of the file as well
// capital case = GLOBAL SCOPED
// lower case = File Global Scoped (save this for a scope video)

var ( // you can use this to declare multiple variables at once when in global scope
	d string  = "A word"
	e bool    = false
	f int     = 9
	g float64 = 14.1567
)

// NOTE: There are A LOT of int and float types to choose from (8, 16, 32, 64 bit, unsigned etc.)

func main() {
	var a int = 10 // declare a variable in go
	fmt.Println(a)

	a = 22
	fmt.Println(a)
	g = float64(f) // type casting (int to float)
	a = int(g)
	println(g, a)

	b := 50 // this will automatically declare a b variable as an integer type
	fmt.Println(b)

	name := "Jack"
	var age int = 20
	println("Hello! My name is ", name, " and I am ", age, "years old.")
	age += 1 // add 1 to age
	println("It's my birthday tomorrow, so I will be ", age)
	println("Now I can say that my name is ", name, " and my age is ", age)
}
