package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines allows for concurrency, this is when you do something
// and while waiting for that something you start doing something
// else and then come back to the thing you were doing

// imagine you have a really slow PC that takes 2 minutes to boot
// you start booting up your PC, then while you wait you go to the
// kitchen to get some food, you then come back to your computer
// now that it's done booting up and start working on it. This example
// is like a goroutine, instead of booting up your PC and waiting for
// it to finish booting up, you went to the kitchen to do something
// else while you waited

// example 1
// imagine below is sending a request to a server
// and that server takes 200 milliseconds to respond
func startPC() {
	fmt.Println("Starting PC")
	// booting takes 3 seconds, imagine this was 2 min
	// like in above example explaination
	time.Sleep(time.Second * 3)
	fmt.Println("PC has started")
}

func makeFood() {
	fmt.Println("Going to make food")
	time.Sleep(time.Second * 3)
	fmt.Println("Food was made")
}

func exmpl1() {
	/* PART 1 */
	startPC()
	makeFood()

	/*
		OUTPUT:
			Starting PC
			PC has started
			Going to make food
			Food was made

		Explain:
			The above took a while to complete (like 6 seconds), because
			we first started our PC, then waited for it to finish booting
			Then after we waited for it to boot we went to go make food
	*/

	/* PART 2 */
	go startPC() // start a go routine
	makeFood()

	/*
		OUTPUT:
			Going to make food
			Starting PC
			Food was made

		Explain:
			The order is not quite right yet, but we started our PC
			and while we waited, we went o make food. Currently we
			did not get to the end of starting the PC, mainly because
			we did not specify that it should complete the routine before
			we end the program
	*/

	/* PART 3 */
	go startPC()
	go makeFood()

	/*
		OUTPUT:
			[nothing]

		Explain:
			Same issue as above, neither of the routines are required to
			finish before exiting the program, so none of them got time
			to finish before starting
	*/

	/* PART 4 */
	go startPC()
	go makeFood()

	time.Sleep(time.Second * 4)

	/*
		OUTPUT:
			Going to make food
			Starting PC
			PC has started
			Food was made

		Explain:
			The order might not be right, but the above did both tasks in
			about 3 seconds (half the time it initially needed!), it is
			becuase we forced the program to keep running for at least
			4 seconds, enough time to start the routines and finish them

			The point is, both of the tasks were done concurrently, we
			started the pc, went to make food, the pc finished startup,
			we finished making the food

			However, take note that using sleep is not a good way to
			handle go routines, instead use waitgroups (next topic)
	*/
}

// example 2
// we can use a wait group instead of sleep to allow goroutines
// to finish
var wg sync.WaitGroup

func exmpl2() {
	/* Part 1 */
	wg.Add(1) // we add 1 item to waitgroup
	go startPC2()

	wg.Add(1)
	go makeFood2()

	wg.Wait() // this will wait for above added go routines to finish

	/*
		OUTPUT:
			Going to make food
			Starting PC
			PC has started
			Food was made

		Explain:
			Now, instead of having to add a sleep, it knows when the
			go routines have finished, and can stop executing on its
			own
	*/

	/* Part 2 */
	wg.Add(1)
	go startPC2()

	wg.Add(1)
	go makeFood2()

	wg.Wait()

	// the below go routine and functions will only start executing after
	// the above wait, this functionality allows you to only execute code
	// if go routines have finished executing
	fmt.Println("Test")
	wg.Add(1)
	go startPC2()
	wg.Wait()

	/*
		OUTPUT:
			Going to make food
			Starting PC
			PC has started
			Food was made

		Explain:
			Now, instead of having to add a sleep, it knows when the
			go routines have finished, and can stop executing on its
			own
	*/
}

func startPC2() {
	fmt.Println("Starting PC")
	time.Sleep(time.Second * 2)
	fmt.Println("PC has started")
	wg.Done() // will remove from waitgroup
}

func makeFood2() {
	fmt.Println("Going to make food")
	time.Sleep(time.Second * 2)
	fmt.Println("Food was made")
	wg.Done()
}

// example 3
var count = 0

func printCount() {
	fmt.Println("Counter is", count)
	wg.Done()
}

func addToCount() {
	count++
	wg.Done()
}

// add in part 2
var lock sync.RWMutex // this will allow us to work with variables

func printCount2() {
	fmt.Println("Counter is", count)
	lock.RUnlock() // removes read lock
	wg.Done()
}

func addToCount2() {
	count++
	lock.Unlock() // removes lock
	wg.Done()
}

func exmpl3() {
	/* Part 1 */
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go printCount()
	// 	go addToCount()
	// }
	// wg.Wait()

	/*
		Output:
			Counter is 1
			Counter is 2
			Counter is 3
			Counter is 4
			Counter is 5
			Counter is 5
			Counter is 6
			Counter is 7
			Counter is 7
			Counter is 8

		Explain:
			Your output may not be exactly the same, but, as you can see,
			we're getting a bunch of weird outputs, instead of getting
			0 to 9, we get the above... This is because 2 go routines
			might be trying to work with the variable at the same time
			which could cause weird things to happen like above
	*/
	count = 0

	for i := 0; i < 10; i++ {
		wg.Add(2)
		// the below will lock any variables being used inside
		// the go routines, this means that other go routines
		// are not allowed to read the variables until it is unlocked
		lock.RLock()
		go printCount2()

		// the below does the sme as above, but does not allow the
		// variables to be mutated/chaned until unlocked
		lock.Lock()
		go addToCount2()
	}
	wg.Wait()

	/*
		Output
			Counter is 0
			Counter is 1
			Counter is 2
			Counter is 3
			Counter is 4
			Counter is 5
			Counter is 6
			Counter is 7
			Counter is 8
			Counter is 9

		Explain:
			We now get the expected outputs, because the go routines wait
			before executing since the variables are locked.
	*/
}

func main() {
	exmpl1()
	exmpl2()
	exmpl3()
}
