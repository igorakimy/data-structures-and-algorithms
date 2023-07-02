package main

// importing fmt package
import (
	"fmt"
)

// gets the power series of integer `a` and returns tuple
// of square of `a` and cube of `a`
func powerSeries(a int) (int, int) {

	square := a * a
	cube := square * a

	return square, cube
}

func main() {

	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square:", square, "Cube:", cube)
}

// Square: 9 Cube: 27
