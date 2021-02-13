package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
	}

	colors["white"] = "#ffffff"

	delete(colors, "green")

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Color: %v - Hex: %v\n", color, hex)
	}
}
