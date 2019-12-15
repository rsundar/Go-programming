package main

import "fmt"

func main() {
	var colors map[string]string

	newColors := make(map[string]string)
	// colors := map[string]string{
	// "red":   "#ff0000",
	// "green": "#74f500",
	// }

	newColors["White"] = "#ffffff"
	fmt.Println(colors)
	fmt.Println(newColors)

	delete(newColors, "White")

	fmt.Println(newColors)
}
