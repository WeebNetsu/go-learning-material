package main

import "fmt"

func main() {
	// basic for loop in go
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip the rest of the loop and continue with the next iteration
		}

		fmt.Println(i)
	}

	// the below is an example where you use a variable
	// that can be used outside of the for loop as well
	var counter int = 5
	for ; counter > 0; counter-- {
		fmt.Println(counter)
	}

	// you are also allowed to change where the counter
	// will be incremented
	for counter < 5 {
		counter++
		fmt.Println(counter)
	}

	for { // this is an infinite loop (almost like a do-while loop)
		counter--

		fmt.Println(counter)

		if counter == 0 {
			break // break out of loop (can be used in a normal for loop too)
		}
	}

	for i := 5; i < 10; i++ {
		for j := 0; j < i; j += 2 { // nested for loop
			fmt.Println(i, j)
		}
	}

	// this method allows you to initalize 2 variables
	// and increment both those variables for the loop
	for i, j := 0, 15; i < 5; i, j = i+1, j*5 {
		fmt.Println(i, j)
	}

	s := [5]int{2, 4, 6, 8, 10}
	for index, value := range s { // to loop through an array type
		fmt.Println(index, value)
	}

	friends := map[string]int{
		"Jack": 18,
		"Nick": 20,
		"Sara": 19,
	}

	for k, v := range friends { // can also do range loop with maps
		println(k, "is", v, "years old")
	}

	for k := range friends { // only get the key/index
		println(k)
	}

	for _, v := range friends { // only get the value
		println(v)
	}
}
