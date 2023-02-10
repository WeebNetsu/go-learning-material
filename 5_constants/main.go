package main

import "fmt"

const ( // can be declared multiple times, just like with variables
	a = iota // iota is used to make an enumirated constant (enum basically) this will be 0
	b = iota // 1
	c = iota // 2 (note: iota is scoped to this constant block, so an iota outside this block will start over)
)

const (
	d = iota             // will be 0
	e                    // iota will be automatically inferred, will be 1
	f                    // 2
	g = 2 << iota        // basically say iota * 2 (you are bitshifting it actually)
	h                    // this will continue the same pattern, so 16 * 2
	i                    // 32 * 64
	j = 1<<iota*iota - 3 // you can do more complex equations (for tutorial, do basic +)
	k
	l
)

func main() {
	// we don't do constants as MY_CONST in Go because, if the first letter of
	// a variable is a capital letter, it will be exported to outside of the file
	const myConst int = 99
	// myConst = 88 // not possible because constant

	// const sin float64 = math.Sin(2.12) // not allowed, since constants are set during compile time
	// and you cannot execute functions during compile time!

	const myName string = "Jack" // constants can contain all data types
	const pi = 3.14159           // constants can also be assigned types automatically

	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g, h, i, j, k, l)
}
