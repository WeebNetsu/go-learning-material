package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Program Start")
	/*
		defer basically keeps code from executing until the function
		is at the end, then right before hitting return, all the defered
		code will execute
	*/
	defer fmt.Println("Program Execution")
	fmt.Println("Program End")

	var x int = 50
	defer fmt.Println(x) // this will print out 50 after the main function executed
	// why? becaise the defer function remembers the exact state the variables
	// were when passed in, so the below variable change will not affect the output
	x *= 2
}

func deferExample() bool {
	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		return false
	}

	/*
		Sometimes you'll be working with a lot of lines of code, and by doing
		so, you can forget to close whatever you're working with (like when
		making an HTTP request), so what you can do is you can tell the resource to close with defer, so all the code that requires it can
		still execute, however, now the closing of it will happen AFTER the
		function finished executing, so you can tell it to close before you
		start reading/working with the resource, ultimately minizing the chances
		you forget to close it later
	*/
	defer res.Body.Close()

	// we close the response body before we read from it, it is possible
	// thanks to defer
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return false
	}

	fmt.Println(robots)

	return true
}
