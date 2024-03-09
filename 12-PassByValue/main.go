package main

import "fmt"

func updateName(name string) string {
	name = "Doni"
	return name
}

func updateMenu(menu map[string]float64) {
	menu["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
    name := "Danuartha"

	updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions

	menu := map[string]float64{
		"pie": 5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}