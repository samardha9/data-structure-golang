package datastructures

import "fmt"

func DemonstrateStrings() {
	str := "Hello, Go!" // Declare a string

	fmt.Println("String demonstration:")
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length: %d\n", len(str))
	fmt.Printf("Character at index 0: %c\n", str[0])
}
