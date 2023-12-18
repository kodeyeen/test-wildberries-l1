package main

import "log"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// тут храним группы
	groups := make(map[int][]float64)

	for _, temp := range temps {
		d := temp / 10.0
		groupKey := int(d) * 10 // определяем ключ группы

		groups[groupKey] = append(groups[groupKey], temp)
	}

	log.Println(groups)
}
