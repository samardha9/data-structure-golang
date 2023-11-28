package main

import "fmt"

func DemonstrateMaps() {
	m := make(map[string]int) // Declare a map with string keys and int values

	// Initialize the map
	m["first"] = 10
	m["second"] = 20
	m["third"] = 30

	fmt.Println("Map demonstration:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
