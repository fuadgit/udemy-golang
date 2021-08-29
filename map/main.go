package main

import "fmt"

func main() {
	//var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// colors["blue"] = "#0000ff"
	// delete(colors, "blue")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code of", color, "is", hex)
	}
}
