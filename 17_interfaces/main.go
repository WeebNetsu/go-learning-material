package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // seed so we can get better random number

	var f Fridge = Fridge{contents: []string{"eggs", "milk"}}
	tv := TV{channel: "NET TV"}
	plug := []Plug{f, tv} // we are implementing the interface instead of the struct

	// using the structs
	/* if f.Pull(2) {
		fmt.Println("Fridge can run with electricity")
	} else {
		fmt.Println("Fridge needs too much power")
	}

	if tv.Pull(2) {
		fmt.Println("TV can run with electricity")
	} else {
		fmt.Println("TV needs too much power")
	} */

	// so here we can run our Pull methon on both our fridge and TV,
	// however, we cannot access their 'content' or 'channel' attributes
	// since this is being run on an interface
	for _, elecDevice := range plug {
		if elecDevice.Pull(2) {
			fmt.Println("Device can run with electricity")
		} else {
			fmt.Println("Device needs too much power")
		}
	}

	// using pointers with interfaces
	myNum := Counter(10)         // create a Counter
	var dec Decrementer = &myNum // assign the memory address to decrementer

	fmt.Println(dec.Decrement()) // can now decrement it
	fmt.Println(dec.Decrement())
	fmt.Println(dec.Decrement())
}

/*
	interfaces are very similar to structs, but instead of
	defining the data that will be stored inside of it, it
	defines the functions

	You don't really specify when something is an interface, it is dermained by
	the struct, if the struct has all of the metods inside of an interface, it
	is automatically of that type of interface

	https://youtu.be/qJKQZKGZgf0
*/
type Plug interface {
	Pull(int) bool // if an applience need to draw power, like a fridge
	// int = how much power is being pulled from the plug
	// bool = if the power tripped because of too much pressure
}

type Fridge struct {
	contents []string
}

type TV struct {
	channel string
}

// chair will not have a plug method, since it does not have any reason
// to be connected to electricity
type Chair struct{}

// because this implements the call method on Fridge, Fridge is now
// part of an interface, if we were to add another function to the Plug interface
// then Fridge would no longer be part of the interface, since it doensn't have
// all of the methods inside the Plug interface (until we add them)
func (fr Fridge) Pull(power int) bool { // allows the fridge to "pull" power from the "plug"
	availiblePower := rand.Intn(10) // random number between (including) 0 and 9

	if power > availiblePower {
		return false
	}

	return true
}

func (tv TV) Pull(power int) bool { // allows the fridge to "pull" power from the "plug"
	availiblePower := rand.Intn(10) // random number between (including) 0 and 9

	if power > availiblePower {
		return false
	}

	return true
}

type Decrementer interface {
	Decrement() int
}

type Counter int // like creatying your own data type (but it's actually a normal integer) - is a struct?

func (ic *Counter) Decrement() int {
	*ic--

	return int(*ic)
}
