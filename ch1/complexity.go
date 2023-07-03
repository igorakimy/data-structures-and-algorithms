package main

import (
	"fmt"
)

func main() {
	var m [10]int
	var k int
	for k = 0; k < 10; k++ {
		m[k] = k + 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}

// Element[0] = 200
// Element[1] = 201
// Element[2] = 202
// Element[3] = 203
// Element[4] = 204
// Element[5] = 205
// Element[6] = 206
// Element[7] = 207
// Element[8] = 208
// Element[9] = 209
