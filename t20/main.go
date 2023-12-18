package main

import (
	"log"
	"strings"
)

func reverseWords(s string) string {
	// сплитим строку на слова
	words := strings.Split(s, " ")
	// слайс для хранения слов в обратном порядке
	reversed := make([]string, len(words))

	// простой реверс слайса
	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversed[j] = words[i]
	}

	// джоиним слова в обратном порядке
	return strings.Join(reversed, " ")
}

func main() {
	log.Println(reverseWords("snow dog sun"))
}
