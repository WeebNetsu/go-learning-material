package main

import "fmt"

func main() {
	sayHey()
	sayHello("Jack", 18)

	math := sum(10, 5)
	fmt.Println(math)

	fmt.Println(calculate(true, 5, 10, 15, 20, 25))
	fmt.Println(calculate(false, 5, 10, 15, 20, 25))

	whatami := "cool"
	fmt.Println(concat("Nick", " I", " am ", whatami))

	division, err := div(19, 8) // will return error if occured
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(division)

	division2, err := div(19, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(division2)

	// youll notice that the name stays the same, it does
	// not change the original name variable
	var name string = "Jack"
	changeNameVer1(name)
	fmt.Println(name)
	/*
		NOTE FOR TUTORIAL
		with the below pointer/reference stuff, just act it off like it's a cool
		go feature with function, we can tell the viewers that it's a pointer in
		another tutorial
	*/
	changeNameVer2(&name) // we now pass in by memory address
	fmt.Println(name)     // will actually change the original variable

	func() { // annonymous function, gets executed on the fly
		fmt.Println("Hello There!")
	}()

	for i := 0; i < 5; i++ {
		func(i int) { // allows you take in arguments
			// passing in a variable is better than getting the data
			// from the larger scope
			fmt.Println(i)
		}(i) // pass in a variable
	}

	var myFunc func(name string) string = func(name string) string { // you can declare a named function like this as well
		fmt.Println("Your name is", name)
		return "Success"
	}

	x := myFunc("Jack") // then use it here
	fmt.Println(x)
}

func sayHey() {
	fmt.Println("Hello World!")
	fmt.Println("I am truly epic!")
	fmt.Println("-- end of function --")
}

func sayHello(name string, age int) {
	fmt.Println("Hello", name, "you are", age, "years old!")
}

func sum(x, y int) int {
	var ans int = x + y
	return ans
}

// this will take all arguments passed in (after argument 1) as a slice
// note you can only have ONE ...dataType in your function parameters
// and it has to be at the END of your arguments
func calculate(add bool, data ...int) int {
	var result int = 0

	for _, value := range data {
		if add {
			result += value
		} else {
			result -= value
		}
	}

	return result
}

// below result will basically already be declared for us
func concat(words ...string) (result string) {
	for _, word := range words {
		result += word
	}

	return // will automatically return result
}

// this allows us to also return an error parameter
func div(num1, num2 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}

	// nil = empty/void/non existent value
	return num1 / num2, nil
}

func changeNameVer1(name string) {
	name = "Nick" // this will not change the original name variable
	fmt.Println(name)
}

func changeNameVer2(name *string) { // it now receives a pointer
	*name = "Nick" // dereference it and change it
}
