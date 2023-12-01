package datastructures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/samar/data-structure-golang/types"
)

func DemonstrateStructs() {
	// Read the JSON file
	jsonFile, err := os.Open("/Users/samar/github/data-structures/data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Read the opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize a struct to hold the data from JSON
	var colorData types.ColorData // Use the ColorData struct from the types package

	// Unmarshal byteArray into the colorData struct
	err = json.Unmarshal(byteValue, &colorData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Struct demonstration with colors from JSON:")
	for i, color := range colorData.Colors {
		fmt.Printf("Index: %d, Color: %s, Category: %s, Type: %s, Hex Code: %s\n", i, color.Color, color.Category, color.Type, color.Code.Hex)
	}
}
