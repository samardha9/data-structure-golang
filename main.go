package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What data structure demonstration do you need?")
	fmt.Println("Select 1 - Array")
	fmt.Println("Select 2 - Slices")
	fmt.Println("Select 3 - Maps")
	fmt.Println("Select 4 - Channels")
	fmt.Println("Select 5 - Strings")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		demonstrateArray()
	case 2:
		demonstrateSlices()
	case 3:
		demonstrateMaps()
	case 4:
		demonstrateChannels()
	case 5:
		demonstrateStrings()
	default:
		fmt.Println("Invalid option selected.")
	}
}

func demonstrateArray() {
	var arr [5]int // Declare an array of 5 integers

	// Initialize the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array demonstration:")
	for i, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
}

func demonstrateSlices() {
	slice := []int{10, 20, 30, 40, 50} // Declare a slice of integers

	fmt.Println("Slice demonstration:")
	for i, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
}

func demonstrateMaps() {
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

func demonstrateChannels() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Message from channel!"
	}()

	fmt.Println("Channel demonstration:")
	msg := <-ch // Receive message from channel
	fmt.Println(msg)
}

func demonstrateStrings() {
	str := "Hello, Go!" // Declare a string

	fmt.Println("String demonstration:")
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length: %d\n", len(str))
	fmt.Printf("Character at index 0: %c\n", str[0])
}
