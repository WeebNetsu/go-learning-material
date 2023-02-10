package main

import (
	"fmt"
	"reflect"
)

type Car struct { // think of it kinda like a pythin dict or json
	// other files will be able to see the Car struct, but none of its
	// content, since the names are start lowercase (owner -> Owner to make it public outside file)
	// basically the OOP of Go
	numDoors   int
	owner      string
	weeklyCost [7]float32
}

func (c Car) getOwner() string { // you can make a function part of a struct as such
	fmt.Println("This car has", c.numDoors, "doors")

	return c.owner
}

type Toyota struct {
	Car  // inheritance, Toyota now contains everything inside of Car
	kind string
}

type Phone struct {
	smart bool
	// below, the "tag" basically says that owner is a required field and can only
	// contain up to 100 characters
	// HOWEVER! It does NOT enforce it, it only provides a piece of text
	owner   string `required,max:"100"`
	android bool
}

func main() {
	stevesCar := Car{
		numDoors: 2,
		owner:    "Steve",
		weeklyCost: [7]float32{
			221.89, 414.29, 100.00, 119.29, 321.11, 40.76, 70.21,
		},
	}

	fmt.Println(stevesCar)
	fmt.Println(stevesCar.numDoors)
	fmt.Println(stevesCar.weeklyCost[1])
	fmt.Println(stevesCar.getOwner()) // you can easily call the function

	jacksToyota := Toyota{}
	jacksToyota.numDoors = 4 // inherited by adding Car struct inside Toyota
	jacksToyota.owner = "Jack"
	jacksToyota.weeklyCost = [7]float32{
		221.89, 414.29, 100.00, 119.29, 321.11, 40.76, 70.21,
	}
	jacksToyota.kind = "Truck"

	/* jacksToyota := Toyota{ // this is the same as above
		Car: Car{
			numDoors: 2,
			owner:    "Steve",
			weeklyCost: [7]float32{
				221.89, 414.29, 100.00, 119.29, 321.11, 40.76, 70.21,
			},
		},
		kind: "Truck",
	} */

	fmt.Println(jacksToyota)

	myPhone := reflect.TypeOf(Phone{})
	field, _ := myPhone.FieldByName("owner")
	fmt.Println(field.Tag) // get the tag
	// the tag is only a string of text, we have to implement
	// something to do with the given info

	// Anonymous struct
	// these are structs that will not exist very long
	jack := struct { // defining the struct
		name string
		age  int
	}{ // setting the struct data
		name: "Jack",
		age:  23,
	}

	fmt.Println(jack)
}
