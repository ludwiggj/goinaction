package main

import (
	"fmt"
)

func createMaps() {
	// Create a map with a key of type string and a value of type int.
	dict1 := make(map[string]int)

	fmt.Println(dict1)

	// Create a map with a key and value of type string.
	// Initialize the map with 2 key/value pairs.
	// This is idiomatic Go code
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict2)

	// Cannot create a map using a slice of strings as the key.
	// dict := map[[]string]int{}

	// Compiler Exception:
	// invalid map key type []string

	// Create a map using a slice of strings as the value - this is fine!
	dict3 := map[int][]string{}
	fmt.Println(dict3)
}

func assigningMapValues() {
	// Create an empty map to store colors and their color codes.
	colors := map[string]string{}

	// Add the Red color code to the map.
	colors["Red"] = "#da1337"

	fmt.Println(colors)
}

func nilMap() {
	// Create a nil map by just declaring the map.
	var colors map[string]string
	fmt.Println(colors)

	// Add the Red color code to the map.
	colors["Red"] = "#da1337"

	// Runtime Error:
	// panic: runtime error: assignment to entry in nil map
}

func retrievingMapValues1(colors map[string]string, key string) {
	// Retrieve the value for the key "Blue".
	value, exists := colors[key]

	// Did this key exist?
	if exists {
		fmt.Println(value)
	} else {
		fmt.Printf("Key %v not present in map\n", key)
	}
}

func retrievingMapValues2(colors map[string]string, key string) {
	// Retrieve the value for the key "Blue".
	value := colors[key]

	// Did this key exist?
	if value != "" {
		fmt.Println(value)
	} else {
		fmt.Printf("Key %v not present in map\n", key)
	}
}

func mapIterateAndDelete() {
	// Create a map of colors and color hex codes.
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	fmt.Println()

	// Remove the key/value pair for the key "Coral".
	delete(colors, "Coral")

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

func passingAMap() {
	// Create a map of colors and color hex codes.
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	fmt.Println()

	// Remove the key/value pair for the key "Coral".
	removeColor(colors, "Coral")

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func main() {
	// createMaps()
	// assigningMapValues()
	// nilMap()

	// colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	// retrievingMapValues1(colors, "Blue")
	// retrievingMapValues1(colors, "Orange")

	// retrievingMapValues2(colors, "Blue")
	// retrievingMapValues2(colors, "Orange")

	// mapIterateAndDelete()
	passingAMap()
}
