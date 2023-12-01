package datastructures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/samar/data-structure-golang/types"
)

// Define a struct to represent a node in the linked list
type ColorNode struct {
	Color types.Color
	Next  *ColorNode
}

// Define a linked list to store colors
type ColorLinkedList struct {
	Head *ColorNode
}

// Function to add a color to the linked list
func (ll *ColorLinkedList) AddColor(color types.Color) {
	newNode := &ColorNode{Color: color, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		currentNode := ll.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

// Function to print the colors in the linked list
func (ll *ColorLinkedList) PrintColors() {
	currentNode := ll.Head
	for currentNode != nil {
		fmt.Printf("Color: %s, Category: %s, Type: %s, Hex Code: %s\n", currentNode.Color.Color, currentNode.Color.Category, currentNode.Color.Type, currentNode.Color.Code.Hex)
		currentNode = currentNode.Next
	}
}

func DemonstrateLinkedList() {
	// Read the JSON file
	jsonFile, err := os.Open("/Users/samar/github/data-structures/data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Read the opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize a linked list to hold the colors
	colorList := &ColorLinkedList{}

	// Unmarshal byteArray into ColorData struct
	var colorData types.ColorData
	err = json.Unmarshal(byteValue, &colorData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Linked List demonstration with colors from JSON:")
	for _, color := range colorData.Colors {
		colorList.AddColor(color)
	}

	// Print the colors in the linked list
	colorList.PrintColors()
}
