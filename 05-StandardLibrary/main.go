package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
    greetings := "Hello there friends!"

	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "th"))
	fmt.Println(strings.Split(greetings, " "))

	ages := []int{50, 23, 18, 19, 34}
	
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 23)
	fmt.Println(index)

	names := []string{"Ida", "Putu", "Sucita", "Danuartha"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Putu"))
}