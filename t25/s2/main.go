// Решение с использованием бесконечного цикла

package main

import "time"

func Sleep(d time.Duration) {
	start := time.Now()

	// цикл займет программу на время, заданное аргументом d
	// пока не прошло d времени с времени старта условие == true
	for time.Since(start) < d {
	}
}

func main() {
	Sleep(5 * time.Second)
}
