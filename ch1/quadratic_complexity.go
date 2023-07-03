package main

import (
	"fmt"
)

func main() {
	var k, l int
	for k = 1; k <= 10; k++ {
		fmt.Println("Multiplication table", k)
		for l = 1; l <= 10; l++ {
			var x = l * k
			fmt.Println(x)
		}
	}
}

// Multiplication table 1
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// Multiplication table 2
// 2
// 4
// 6
// 8
// 10
// 12
// 14
// 16
// 18
// 20
// Multiplication table 3
// 3
// 6
// 9
// 12
// 15
// 18
// 21
// 24
// 27
// 30
// Multiplication table 4
// 4
// 8
// 12
// 16
// 20
// 24
// 28
// 32
// 36
// 40
// Multiplication table 5
// 5
// 10
// 15
// 20
// 25
// 30
// 35
// 40
// 45
// 50
// Multiplication table 6
// 6
// 12
// 18
// 24
// 30
// 36
// 42
// 48
// 54
// 60
// Multiplication table 7
// 7
// 14
// 21
// 28
// 35
// 42
// 49
// 56
// 63
// 70
// Multiplication table 8
// 8
// 16
// 24
// 32
// 40
// 48
// 56
// 64
// 72
// 80
// Multiplication table 9
// 9
// 18
// 27
// 36
// 45
// 54
// 63
// 72
// 81
// 90
// Multiplication table 10
// 10
// 20
// 30
// 40
// 50
// 60
// 70
// 80
// 90
// 100
