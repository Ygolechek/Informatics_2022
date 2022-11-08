package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	var a float64 = 0.4
	var b float64 = 0.8

	return ((math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a / b))) * math.Pow(a*b, 1.0/3.0)
}

func main() {
	fmt.Printf("=======\nЗадача A\n=======\n")
	startValueForX := 3.2
	endValueForX := 6.2
	step := 0.6
	for x := startValueForX; x <= endValueForX; x += step {
		fmt.Printf("y(%.2g) = %.13g \n", x, f(x))
	}

	fmt.Printf("=======\nЗадача B\n=======\n")
	variableValues := [5]float64{4.48, 3.56, 2.78, 5.28, 3.21} // array
	for index, element := range variableValues {
		fmt.Printf("y(x%d) = y(%g) = %.13g \n", index+1, element, f(element))
	}
}
