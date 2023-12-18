package main

import (
	"log"
)

var numbers = [...]int{2, 4, 6, 8, 10}

const (
	taskCnt   = len(numbers)
	workerCnt = 3
)

func main() {
	tasks := make(chan int, taskCnt)
	results := make(chan int, taskCnt)

	// отправляем в планировщик горутины воркера
	for i := 0; i < workerCnt; i++ {
		go worker(tasks, results)
	}

	// заполняем канал tasks
	for _, n := range numbers {
		tasks <- n
	}

	// закрываем канал, т.к. больше не будем туда записывать ничего
	close(tasks)

	// проверяем результаты
	for i := 0; i < taskCnt; i++ {
		log.Println(<-results)
	}

	// закрываем канал, т.к. больше не будем ничего считывать из него
	close(results)
}

func worker(tasks <-chan int, results chan<- int) {
	for t := range tasks {
		// делаем какую-то работу и кладем результат в results
		results <- square(t)
	}
}

func square(n int) int {
	return n * n
}
