// если порядок элементов НЕ имеет значения

package main

import "log"

func main() {
	example := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	log.Println("До удаления:   ", example)

	example = remove(example, 4) // удаляем 5-ку

	log.Println("После удаления:", example)
}

func remove(slice []int, i int) []int {
	length := len(slice)

	// заменяем удаляемый элемент последним
	slice[i] = slice[length-1]

	// возвращаем слайс исключая последний элемент (т.к. его "переместили")
	return slice[:length-1]
}
