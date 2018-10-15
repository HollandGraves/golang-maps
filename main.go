package main

// IMPORTS GO HERE

import "fmt"

// STRUCTS GO HERE

// MAPS GO HERE

// MAIN GOES HERE

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	printMap(colors)
}

// USER DEFINED FUNCTIONS GO HERE
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
