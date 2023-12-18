package main

import "log"

func BinarySearch(needle int, haystack []int) int {
	low := 0                  // индекс начала
	high := len(haystack) - 1 // индекс конца

	// пока значения не "схлопнутся"
	for low <= high {
		mid := low + (high-low)/2 // берем середину
		guess := haystack[mid]    // элемент посередине

		if guess == needle {
			return mid
		} else if guess > needle {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	numbers := []int{0, 1, 4, 9, 16, 25, 36, 49, 64}
	index := BinarySearch(16, numbers)

	if index == -1 {
		log.Println("Item not found")
	} else {
		log.Printf("Found item at index: %d", index)
	}
}
