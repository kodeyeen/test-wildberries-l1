package main

import (
	"log"
	"time"
)

func main() {
	// завершение через специальный канал
	c := make(chan struct{})

	go routine(c)

	time.Sleep(5 * time.Second)

	c <- struct{}{} // посылаем сигнал о завершении горутины
}

func routine(c <-chan struct{}) {
	for {
		select {
		case <-c: // return как только придет сигнал
			log.Println("Finished")
			return
		default:
			log.Println("Doing something...")
		}
	}
}
