package main

import "fmt"

func main() {
	var age int = 9

	switch age {
	case 18:
		fmt.Println("You are just old enough to enter!")
	case 21, 30, 50, 100:
		fmt.Println("You've reached an age milestone! Enjoy your day!")
	default:
		fmt.Println("You may not enter")
	}

	switch { // this allows for boolean comparison
	case age < 10:
		fmt.Println("Just... No... NO KIDS ALLOWED.")

		// this will let the following case statement execute
		// (if it's true or not), don't use it too often
		fallthrough
	case age < 18, age > 100:
		fmt.Println("You are too young or too old. You may not enter")
	default:
		fmt.Println("You may enter!")
	}

	// just like with an if statement, we set the variable, then we check
	// the variable (and is scopped to this switch block)
	switch x := 9 * 5; x {
	case 45:
		fmt.Println("Is 45")
	}
}
