package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Writing to a file
	// create a file
	file, err := os.Create("./createdFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, "Hello! This is text being written into a file")

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of text written to file is:", length)

	file.Close() // close the file after using it

	// reading from a file
	// note: data is an array of aschii values
	data, err := ioutil.ReadFile("./readme.txt")

	if err != nil {
		panic(err)
	}

	// by wrapping data in a string(), you convert the bytes to string
	fmt.Println(string(data))
}
