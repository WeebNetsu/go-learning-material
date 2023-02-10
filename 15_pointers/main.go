package main

import "fmt"

type Person struct {
	name string
}

func main() {
	var a int = 10
	var b int = 20
	a = 46 // only 'a' will change, b will stay the same

	fmt.Println(a, b)

	// & returns the memory address of a variable
	var c *int = &a // create a pointer that holds the memory location of a

	fmt.Println(&a, c) // memory address of the 'a' variable
	fmt.Println(*c)    // dereference the variable to get the value in 'a'
	a = 22             // since 'c' just points to 'a', c will also change if we dereference it
	fmt.Println(*c)
	*c = 19 // will also change 'a'
	fmt.Println(a)

	var jack *Person // allowed to create a struct pointer
	jack = &Person{name: "Jack"}
	fmt.Println(jack)

	var nick *Person
	fmt.Println(nick)  // by default this would be a nil pointer
	nick = new(Person) // can also use 'new' instead of '&' here
	fmt.Println(nick)
	(*nick).name = "Nick" // can then set values as such
	fmt.Println((*nick).name)
	// HOWEVER, Gos compilier is smart enough to know that we dont
	// want to access/change the .name on the memory address, but on
	// the value inside the address, so the below is allowed
	nick.name = "Nickole"
	fmt.Println(nick.name)
}
