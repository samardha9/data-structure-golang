package datastructures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/samar/data-structure-golang/types"
)

func DemonstrateArray() {
	// Read the JSON file
	jsonFile, err := os.Open("/Users/samar/github/data-structures/data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Read the opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize our ColorData struct
	var colorData types.ColorData

	// Unmarshal byteArray into 'colorData'
	err = json.Unmarshal(byteValue, &colorData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Display the colors from the JSON file
	fmt.Println("Array demonstration with colors from JSON:")
	for i, color := range colorData.Colors {
		fmt.Printf("Index: %d, Color: %s, Hex Code: %s\n", i, color.Color, color.Code.Hex)
	}
}
