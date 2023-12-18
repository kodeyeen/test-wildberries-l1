package main

import "log"

func someFunc() string {
	// выделяем столько памяти, сколько нам нужно
	// вместо строки длиной 1024 символа, хватит и 128 (2**7)
	v := createHugeString(1 << 7)
	justString := v[:100]

	// не используем глобальную переменную, а возвращаем значение из функции
	return justString
}

func createHugeString(size int) string {
	bytes := make([]rune, size)

	return string(bytes)
}

func main() {
	str := someFunc()

	log.Println(str)
}
