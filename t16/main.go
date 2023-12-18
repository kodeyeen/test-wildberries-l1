package main

import (
	"log"
	"math/rand"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		// если текущий элемент < опорного
		if arr[j] < pivot {
			// то перемещаем его в начало, а начальный на его место
			arr[i], arr[j] = arr[j], arr[i]
			// сдвигаем нижний индекс
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}

	return arr
}

func main() {
	unsorted := make([]int, 15)

	for i := 0; i < 15; i++ {
		unsorted[i] = rand.Intn(100)
	}

	sorted := quickSort(unsorted, 0, len(unsorted)-1)

	log.Println(sorted)
}
