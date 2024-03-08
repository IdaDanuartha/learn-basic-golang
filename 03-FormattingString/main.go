package main

import "fmt"

func main() {
	name := "Danuartha"
	age := 19
	score := 250.55

	// print
	fmt.Print("Hello, ")
	fmt.Print("world \n\n")

	// println
	fmt.Println("Hello my name is", name, "and my age is", age)

	// printf (formatted strings) | %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", score)
	fmt.Printf("you scored %0.1f points! \n\n", score)

	// sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println(str)
}