package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#f80000",
		"green": "#f4010",
		"white": "#	ffffff",
	}

	// this is the way to initialize a map if its not created by the previous syntax
	// colors := make(map[string]string)

	// For add and deletion
	// colors["Chivo"] = "#ff0000"
	// delete(colors, "Chivo")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
