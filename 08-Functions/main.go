package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Good bye %v \n", name)
}

func cycleNames(name []string, cb func(string)) {
	for _, v := range name {
		cb(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
    sayGreeting("Danuartha")
    sayBye("Dono")

	cycleNames([]string{"Ida", "Putu", "Danuartha"}, sayGreeting)

	area1 := circleArea(7.0)
	area2 := circleArea(10.0)

	fmt.Printf("area1 = %0.2f and area2 = %0.2f", area1, area2)
}