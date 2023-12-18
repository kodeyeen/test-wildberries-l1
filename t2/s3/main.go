package main

import (
	"log"
)

var numbers = [...]int{2, 4, 6, 8, 10}

func main() {
	for _, n := range numbers {
		go func(n int) {
			log.Println(square(n))
		}(n)
	}
}

func square(n int) int {
	return n * n
}
