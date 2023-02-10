package main

import "fmt"

func main() {
	friends := [3]string{"Jack", "Nick", "Luke"} // an array that can store 3 string values
	fmt.Println(friends)

	coolNumbers := [...]int{1, 7, 9, 0, 8, 17} // this array will be made just big enough to hold the elements put in it
	fmt.Println(coolNumbers[1])                // display item at index 1 (location 2)

	friends[0] = "Patric" // reassign index in array
	fmt.Println(friends)

	fmt.Println(len(coolNumbers)) // returns number of items in array

	var enemies [3]string // set an empty array variable
	enemies[0] = "Lucas"  // assign values later
	enemies[1] = "Jack"
	enemies[2] = "Simpson"
	fmt.Println(enemies)

	var xoxo [3][3]int = [3][3]int{ // multi dimentional array
		{1, 2, 0},
		{1, 0, 2},
		{2, 1, 2},
	}

	fmt.Println(xoxo[1][0])
}
