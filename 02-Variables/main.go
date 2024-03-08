package main

import "fmt"

func main() {
    // strings
	var name1 string = "Danuartha"
	var name2 = "Ida Putu"
	var name3 string

	fmt.Println(name1, name2, name3)

	name3 = "Sucita"

	fmt.Println(name1, name2, name3)

	name4 := "Alex"

	fmt.Println(name4)

	// integers
	var number1 int = 10
	var number2 = 20
	var number3 int

	fmt.Println(number1, number2, number3)

	// bits & memory
	var bit1 int8 = 25
    var bit2 int8 = -128
	var bit3 uint16 =  256

    fmt.Println(bit1, bit2, bit3)

	// float
	var float1 float32 = 45.90
    var float2 float64 = 459469854854.45

    fmt.Println(float1, float2)
}