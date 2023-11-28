package datastructures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/samar/data-structure-golang/types"
)

func DemonstrateChannels() {
	// Read and parse the JSON file
	colorData, err := readColorData("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a channel for Color type
	ch := make(chan types.Color)

	go func() {
		// Send each color to the channel with a delay
		for _, color := range colorData.Colors {
			time.Sleep(1 * time.Second)
			ch <- color
		}
		close(ch) // Close the channel after sending all colors
	}()

	fmt.Println("Channel demonstration with colors from JSON:")
	// Receive and print each color from the channel
	for color := range ch {
		fmt.Printf("Color: %s, Hex Code: %s\n", color.Color, color.Code.Hex)
	}
}

// readColorData reads and parses the JSON file
func readColorData(filename string) (types.ColorData, error) {
	var colorData types.ColorData

	jsonFile, err := os.Open(filename)
	if err != nil {
		return colorData, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &colorData)
	if err != nil {
		return colorData, err
	}

	return colorData, nil
}
