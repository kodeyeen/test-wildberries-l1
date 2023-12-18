// Решение только для строк содержащих символы английского алфавита

package main

import "log"

func areAllCharsUnique(s string) bool {
	var seen [26]bool

	for _, c := range s {
		// переводим символ в lower case и маппим код символа на индекс в слайсе
		i := (c | 0b100000) - 'a'

		// если уже встречали этот символ
		if seen[i] {
			return false
		}

		seen[i] = true
	}

	return true
}

func main() {
	log.Println(areAllCharsUnique("abcd"))
	log.Println(areAllCharsUnique("abCdefAaf"))
	log.Println(areAllCharsUnique("aabcd"))
}
