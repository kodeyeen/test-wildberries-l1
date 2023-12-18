package main

import "fmt"

func main() {
	a, b := 11111, 99999

	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("After swapping: a = %d, b = %d\n", a, b)
}
