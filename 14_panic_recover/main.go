package main

import "fmt"

func main() {
	// x, y := 10, 0
	// fmt.Println(x / y) // will throw a panic, since you cant divide by 0

	// fmt.Println("Program Start")
	// // cusom panic
	// panic("Some error happened") // when a panic happens, the program terminates
	// fmt.Println("Program End")

	/*
		NOTE:
			Most of the time in Go, a program will not panic, instead it will
			return an error message, then we as the user can decide when to
			throw a panic and when to let the error go
	*/

	// NOTE: defer will still execute, even if panic is thrown
	// defer fmt.Println("Hello World")
	// panic("Oof I crashed")

	fmt.Println("I have started")
	masterPanic(true) // if true, panic
	// masterPanic(false)
	// even if it panic, our program will still continue thanks to the
	// defer we added
	fmt.Println("I have ended")
}

func masterPanic(pls bool) {
	fmt.Println("I am going to panic :)")

	defer func() { // define a defer above the panic
		if err := recover(); err != nil { // if an error was thrown, use recover to get the error
			fmt.Println("ERROR!", err)
		}
	}()

	if pls {
		panic("I'm in panic mode!!!") // if in panic, it will not run the print at the end of the func
	}

	fmt.Println("Panic has completed")
}
