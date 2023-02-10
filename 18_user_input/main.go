package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter your name: ")

	// so we're telling the program to get input from the
	// standard input (stdin) (the terminal)
	// this is only the reader and won't get the input itself
	reader := bufio.NewReader(os.Stdin)

	// the below will get the user input, and stop receiving
	// user input on a new line (when they press enter), thanks
	// to the newline character we entered (\n)
	// if we entered 'c' instead, then once we enter a 'c' into
	// our input, it won't actaully read any of the data entered
	// after it, eg. We enter: "netsucsd"; We get "netsuc"
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error happened while trying to get your name...")
		return
	}

	name = strings.TrimSpace(name) // remove all unused whitespace from a string, like the new line we get with the input

	fmt.Println("Your name is", name)
}
