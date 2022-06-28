package main

import (
	"fmt"
)

func main() {

	//Maps are like dictionary in C# or objects in Javascript, they are key-value pairs
	// All keys have to be of the same type

	//var colors map[string]string
	//colors := make(map[string]string)

	// colors["white"] = "#fff"
	// delete(colors, "white")

	//------  map[type of map]value
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4bf745",
		"white": "#fff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
