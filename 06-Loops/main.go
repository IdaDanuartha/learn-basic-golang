package main

import "fmt"

func main() {
    // loops
	x := 0
	for x < 5 {
		fmt.Println("value of x is :", x)
		x++
	}

	fmt.Println()

	for y := 0; y < 5; y++ {
		fmt.Println("value of y is :", y)
	}

	names := []string{"Ida", "Putu", "Sucita", "Danuartha"}

	fmt.Println()

	for z := 0; z < len(names); z++ {
		fmt.Println(names[z])
	}

	fmt.Println()

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}
}