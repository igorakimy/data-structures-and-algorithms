package main

import "fmt"

func main() {
	var languages = map[int]string{
		3: "English",
		4: "French",
		5: "Spanish",
	}

	var products = make(map[int]string)
	products[1] = "chair"
	products[2] = "table"

	var i int
	var value string
	for i, value = range languages {
		fmt.Println("language", i, ":", value)
	}
	fmt.Println("product with key 2", products[2])

	delete(products, 1)
	fmt.Println("products", products)
}

// language 3 : English
// language 4 : French
// language 5 : Spanish
// product with key 2 table
// products map[2:table]
