/* ! BEFORE TUTORIAL!!! READ whathappens.txt !!! */

package custom // can't be main

import "fmt"

// make them start with capital letters to make them globally usable
func SayHello(name string) {
	fmt.Println("Hello", name)
}

func Add(x, y int) int {
	return x + y
}
