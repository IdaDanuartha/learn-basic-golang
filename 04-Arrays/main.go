package main

import "fmt"

func main() {
	// arrays
    // var ages [3]int = [3]int{20,25,30}
	var ages = [3]int{20,30,40}
	fmt.Println(ages, len(ages))

	names := [4]string{"ida", "putu", "sucita", "danuartha"}
	names[1] = "bagus"
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 200, 300}
	scores[2] = 500
	scores = append(scores, 1000)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:2]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "anang")

	fmt.Println(rangeOne)
}