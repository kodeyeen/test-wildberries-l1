package main

import (
	"fmt"
)

func reverseString(input string) string {
	// преобразуем строку в слайс рунов
	runes := []rune(input)
	// слайс рун для перевернутой строки
	reversed := make([]rune, len(runes))

	// заполняем reversed с конца
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		reversed[j] = runes[i]
	}

	return string(reversed)
}

func main() {
	input := "日本語 авылдф"
	reversed := reverseString(input)

	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
