// если порядок элементов имеет значение, похож на метод с append

package main

import "log"

func main() {
	example := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	log.Println("До удаления:   ", example)

	example = remove(example, 4) // удаляем 5-ку

	log.Println("После удаления:", example)
}

func remove(slice []int, i int) []int {
	// заменяем элементы начиная с удаляемого элементами после удаляемого
	copy(slice[i:], slice[i+1:])

	// исключаем последний элемент
	return slice[:len(slice)-1]
}
