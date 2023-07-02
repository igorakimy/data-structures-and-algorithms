package main

import (
	"fmt"
)

func main() {
	var m [10]int
	var k int
	for k = 0; k < len(m); k++ {
		m[k] = k * 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}

// Element[0] = 0
// Element[1] = 200
// Element[2] = 400
// Element[3] = 600
// Element[4] = 800
// Element[5] = 1000
// Element[6] = 1200
// Element[7] = 1400
// Element[8] = 1600
// Element[9] = 1800
