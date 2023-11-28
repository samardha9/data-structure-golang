package main

import (
	"fmt"

	"github.com/samar/data-structure-golang/datastructures"
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
		datastructures.DemonstrateArray()
	case 2:
		datastructures.DemonstrateSlices()
	case 3:
		datastructures.DemonstrateMaps()
	case 4:
		datastructures.DemonstrateChannels()
	case 5:
		datastructures.DemonstrateStrings()
	default:
		fmt.Println("Invalid option selected.")
	}
}
