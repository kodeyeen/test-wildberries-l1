package main

import (
	"log"
)

func main() {
	var test interface{} = 5

	// type switch
	switch test.(type) {
	case int:
		log.Println("int")
	case string:
		log.Println("string")
	case bool:
		log.Println("bool")
	case chan int:
		log.Println("chan int")
	case chan string:
		log.Println("chan string")
	case chan bool:
		log.Println("chan bool")
	default:
		log.Println("unknown type")
	}
}
