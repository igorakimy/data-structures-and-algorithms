package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 4, 5, 6}

	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println("printing elements", arr[i])
	}

	for _, value := range arr {
		fmt.Println("range", value)
	}

	for _, value := range arr {
		fmt.Println("blank range", value)
	}
}

// printing elements 1
// printing elements 2
// printing elements 4
// printing elements 5
// printing elements 6
// range 1
// range 2
// range 4
// range 5
// range 6
// blank range 1
// blank range 2
// blank range 4
// blank range 5
// blank range 6
