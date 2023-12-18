package main

import (
	"context"
	"log"
	"time"
)

func main() {
	N := 6 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), N)
	defer cancel()

	c := make(chan int)

	go write(c) // горутина для записи в канал
	go read(c)  // горутина для чтения из канала

	<-ctx.Done() // тут блокируемся, но когда прилетит пустая структура (через N секунд) пойдем дальше и программа завершится
	// time.Sleep(N) // либо просто Sleep на N секунд
}

func write(c chan<- int) {
	// после записи всех чисел закрываем канал
	defer close(c)

	for i := 0; i < 100; i++ {
		// записываем в канал каждую секунду
		time.Sleep(1 * time.Second)
		c <- i
	}
}

func read(c <-chan int) {
	for n := range c {
		// читаем из канала и просто выводим
		log.Println(n)
	}
}
