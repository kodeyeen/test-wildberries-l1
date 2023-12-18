package main

import (
	"log"
	"strings"
)

func areAllCharsUnique(s string) bool {
	// если бы строка была ASCII-only то можно было бы избежать этого вызова
	// и преобразовывать в lower case с помощью побитовой операци (сделал так в решении №2)
	s = strings.ToLower(s)
	// отслеживаем, какие символы были просмотрены
	seen := make(map[rune]bool)

	for _, c := range s {
		// если этот символ уже был встречен однажды
		if seen[c] {
			return false
		}

		// иначе сохраняем
		seen[c] = true
	}

	// если в цикле не произошел return, то значит, все символы уникальны
	return true
}

func main() {
	log.Println(areAllCharsUnique("abcd"))
	log.Println(areAllCharsUnique("abCdefAaf"))
	log.Println(areAllCharsUnique("aabcd"))
}
