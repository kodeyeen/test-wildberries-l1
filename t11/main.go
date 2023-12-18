package main

import "log"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	log.Println(intersection(slice1, slice2))
}

func intersection(a, b []int) []int {
	// сюда сохраняем уникальные значения из первого среза
	set := make(map[int]struct{})
	result := []int{}

	// заполняем map уникальными значениями из первого среза
	for _, v := range a {
		set[v] = struct{}{}
	}

	// проверяем второй срез на пересечение с первым
	for _, v := range b {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}

	return result
}
