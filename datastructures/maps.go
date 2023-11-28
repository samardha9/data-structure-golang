package datastructures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/samar/data-structure-golang/types"
)

func DemonstrateMaps() {
	// Read the JSON file
	jsonFile, err := os.Open("/Users/samar/github/data-structures/data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Read the opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize a map with string keys to hold the data from JSON
	var colorData types.ColorData

	// Unmarshal byteArray into the map
	err = json.Unmarshal(byteValue, &colorData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Map demonstration with colors from JSON:")
	for key, color := range colorData.Colors {
		fmt.Printf("Key: %d, Color: %s, Hex Code: %s\n", key, color.Color, color.Code.Hex)

	}
}
