package main

import "fmt"

func main() {
	var arr = []int{5, 6, 7, 8, 9}
	var slice1 = arr[:3]
	fmt.Println("slice1", slice1)
	var slice2 = arr[1:5]
	fmt.Println("slice2", slice2)
	var slice3 = append(slice2, 12)
	fmt.Println("slice3", slice3, cap(slice3))
}

// slice1 [5 6 7]
// slice2 [6 7 8 9]
// slice3 [6 7 8 9 12]
