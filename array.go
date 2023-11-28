package main

import "fmt"

func DemonstrateArray() {
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
