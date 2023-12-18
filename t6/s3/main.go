package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	tasks := make(chan int)

	wg.Add(2)
	go generator(tasks, wg)
	go routine(tasks, wg)

	wg.Wait()
}

func generator(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		n := rand.Intn(100)

		c <- n

		// в какой-то момент решаем завершить работу и закрыть канал
		// закрытие канала скажется на другой горутине (она остановится)
		if n == 3 {
			close(c)
			break
		}
	}
}

func routine(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range tasks {
		log.Println("Doing something...", t)
	}
}
