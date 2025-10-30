package main

import "fmt"

func main() {
	colors := map[string]string{
		"orange": "#ffa500",
		"white":  "#ffffff",
		"green":  "#00ff00",
	}

	aMap := make(map[string]string)
	fmt.Printf("aMap: %v\n", aMap)

	colors["violet"] = "#8f00ff"

	delete(colors, "violet")
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}
