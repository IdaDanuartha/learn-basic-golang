package main

import (
	"fmt"
)

func main() {
    menu := map[string]float64{
		"soup": 4.90,
		"pie": 7.99,
		"salad": 5.99,
	}

	fmt.Println(menu["pie"])

	// looping keys
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}
}