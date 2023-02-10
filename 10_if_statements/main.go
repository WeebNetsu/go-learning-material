package main

import "fmt"

func main() {
	var age int = 12
	var birthday bool = false

	if age > 17 {
		fmt.Println("You may enter")

		if birthday && age == 21 {
			fmt.Println("It's your 21st! All your drinks are on us!")
		}

		if !birthday {
			fmt.Println("Have a nice day, it is not your birthday!")
		}

		// other operators: != <
	} else if age >= 10 || age <= 17 {
		fmt.Println("Teenagers are not allowed near the bar! BEGONE!")
	} else {
		fmt.Println("You are too young and may not enter!")

		if birthday {
			fmt.Println("PizzaHut is 2 streets down")
		}
	}

	friends := map[string]int{
		"Jack":  18,
		"Nick":  24,
		"Sally": 19,
	}

	// below will set age and exists variables, then check exists
	// age and exists are also now scopped to this if statement
	// and can't be used outside of it
	if age, exists := friends["Nick"]; exists {
		fmt.Println("Nick does exist and is", age, "years old")
	}
}
