package main

import (
	"log"
	"sync"
)

var numbers = [...]int{2, 4, 6, 8, 10}

const (
	taskCnt   = len(numbers)
	workerCnt = 3
)

func main() {
	tasks := make(chan int, taskCnt)
	results := make(chan int, taskCnt)

	var wg sync.WaitGroup

	// отправляем в планировщик горутины воркера
	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		go worker(tasks, results, &wg)
	}

	// заполняем канал tasks
	for _, n := range numbers {
		tasks <- n
	}

	close(tasks)

	// блокируемся и ждем, пока воркеры завершат свои работы
	wg.Wait()

	close(results)

	// проверяем результаты
	for r := range results {
		log.Println(r)
	}
}

func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range tasks {
		// делаем какую-то работу и кладем результат в results
		results <- square(t)
	}
}

func square(n int) int {
	return n * n
}
