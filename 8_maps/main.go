package main

import "fmt"

func main() {
	friends := map[string]int{
		"Jack": 18, // maps are like key/value pairs
		"Nick": 23,
		"Luke": 19,
	}

	fmt.Println(friends)
	fmt.Println(friends["Jack"]) // returns that keys value
	friends["Nick"] = 17         // you can modify values in map
	friends["Nirelle"] = 21      // you can add new elements to maps
	delete(friends, "Jack")      // you can delete elements from maps
	fmt.Println(friends)
	value, exists := friends["Lily"] // value will default to 0
	fmt.Println(value, exists)       // if the key does not exist, exists will become false
	value, exists = friends["Nick"]
	fmt.Println(value, exists)
	fmt.Println(len(friends)) // returns number of keys in map

	// NOTE: maps will not always return the values in order when you print the whole map out!

	worstEnemies := make(map[int]string) // if you need a map, but don't have the specific values at the time of creation

	worstEnemies = map[int]string{ // key can be int, bool, string, array etc.
		1: "Cole",
		2: "Mike",
		3: "Zera",
		4: "kali",
	}

	fmt.Println(worstEnemies)

	enemiesCopy := worstEnemies
	enemiesCopy[3] = "Zharani" // NOTE: This will change value in both original and copy of map (delete will also change both)
	fmt.Println(worstEnemies)
	fmt.Println(enemiesCopy)
}
