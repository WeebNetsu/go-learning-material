package main

import "fmt"

func main() {
	friends := []string{"Jack", "Nick", "Caleb", "Mike", "Lucas", "Nigel"} // a slice, no size limit
	fmt.Println(friends)
	fmt.Println(len(friends))                     // get slice size
	friends = append(friends, "Kaali", "William") // add items to slice
	// the ... is like in javascript, it will convert {1, 2} into 1, 2
	friends = append(friends, []string{"Ethan", "Redric"}...)
	fmt.Println(friends)

	a := friends  // NOTE: unlike arrays, if you do this, you copy the original slice
	a[1] = "Loki" // will change in both slices
	fmt.Println(friends)
	fmt.Println(a)

	// note below slicing can be used in both Arrays and slices
	fmt.Println(a[:])   // return all elements in slice
	fmt.Println(a[2:])  // return all elements in slice from index 2 onwards
	fmt.Println(a[:4])  // return all elements in slice from before index 4 (excluding index 4)
	fmt.Println(a[2:4]) // return all elements in slice between index 2 and 4 (excluding index 4)

	b := make([]int, 3) // create a slice with 3 items in it (0 by default)
	fmt.Println(b)

	c := make([]int, 3, 100) // we are creating a slice that has 3 items in it, but can store up to 100
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c)) // capacity is to disply max size of slice

	d := []int{1, 2, 3, 4, 5}
	e := d[1:]                   // remove first element in slice
	f := d[:len(d)-1]            // remove last element in slice
	g := append(d[:2], d[3:]...) // remove element from middle of slice
	// NOTE: g modifies the original slice!!!!
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
