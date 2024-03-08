package main

import "fmt"

func main() {
	age := 40

    if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is no less than 40")
	}

	names := []string{"Ida", "Putu", "Sucita", "Danuartha"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Println("the value at pos %v is %v \n", index, value)
	}
}