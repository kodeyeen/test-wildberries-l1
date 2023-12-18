package main

import "log"

func createSet(strings []string) []string {
	// отслеживаем элементы как ключи мапа
	seen := make(map[string]struct{})
	var set []string

	for _, s := range strings {
		// если строки нет в мапе
		if _, exists := seen[s]; !exists {
			// добавляем в мап и в итоговый слайс
			seen[s] = struct{}{}
			set = append(set, s)
		}
	}

	return set
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := createSet(sequence)

	log.Println(set)
}
