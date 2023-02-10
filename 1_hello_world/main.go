package main // the main package is the main file in your go app

import "fmt" // import fmt to print output

func main() { // entry point of application like in c++
	fmt.Println("Hello World!")
	fmt.Print("Hello World!\n")                    // same as above but without an endlin
	fmt.Printf("%s is %d years old\n", "Nick", 19) // printf for those who like it
	println("Hello there!")                        // also works, but might be removed in the future
	print("Hello there!\n")                        // same as above
	fmt.Println("Hello\tWorld\vI am\ncool")        // you can use escape characters
}

// run your code: go run main.go

// compile: go build github.com/weebnetsu/1_hello_world
// install (compiled) application to bin folder: go install github.com/weebnetsu/1_hello_world
