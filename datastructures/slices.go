package datastructures

import "fmt"

func DemonstrateSlices() {
	slice := []int{10, 20, 30, 40, 50} // Declare a slice of integers

	fmt.Println("Slice demonstration:")
	for i, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
}
