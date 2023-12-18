package main

import "fmt"

func main() {
	a, b := 11111, 99999

	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)

	// арифметический трюк
	a = a + b // overflow risk
	b = a - b
	a = a - b

	fmt.Printf("After swapping: a = %d, b = %d\n", a, b)
}
