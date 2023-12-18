package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	workerCnt := 4
	bufSize := 10
	c := make(chan int, bufSize)

	// закидываем воркеры в планировщик
	for i := 0; i < workerCnt; i++ {
		go worker(c)
	}

	// также генератор, который будет записывать что-то в наш канал
	go generator(c)

	// блокируемся, но следим за syscall, чтобы выйти при нажатии Ctrl + C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // отслеживаем

	<-quit // завершаем работу, когда прилетит значение
}

func generator(c chan<- int) {
	// бесконечно что-то записываем
	for {
		time.Sleep(500 * time.Millisecond)
		c <- rand.Intn(100)
	}
}

func worker(c <-chan int) {
	// бесконечно обрабатываем (выводим) значения в канале
	for {
		log.Println(<-c)
	}
}
