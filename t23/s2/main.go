// если порядок элементов имеет значение (неэффективно)

package main

import "log"

func main() {
	example := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	log.Println("До удаления:   ", example)

	example = remove(example, 4) // удаляем 5-ку

	log.Println("После удаления:", example)
}

func remove(slice []int, i int) []int {
	// слайс до удаляемого элемента не включая его + слайс с элемента после удаляемого и до конца
	return append(slice[:i], slice[i+1:]...)
}
