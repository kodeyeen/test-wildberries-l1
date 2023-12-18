package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	numbers := []int{2, 4, 6, 8, 10}
	// сюда закидываем числа
	tasks := make(chan int)
	// сюда закидываем их квадраты
	results := make(chan int)

	// первая горутина отправляет число для возведения в квадрат в канал tasks
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(tasks)

		for _, n := range numbers {
			// после записи числа произойдет блокировка и запустится другая горутина
			// блокировка будет до тех пор, пока канал не освободится,
			// чтобы мы могли записать следующее число
			tasks <- n
		}
	}()

	// вторая горутина читает число из tasks, возводит в квадрат и записывает в results
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(results)

		// освобождает канал tasks, беря оттуда число
		// возводит его в квадрат и отправляет в results (если можно)
		for t := range tasks {
			results <- square(t)
		}
	}()

	// третья горутина берем из results квадрат числа и прибавляет к сумме
	wg.Add(1)
	go func() {
		sum := 0

		defer wg.Done()

		// освобождает канал results
		for s := range results {
			sum += s
		}

		log.Println("SUM", sum)
	}()

	// ждем завершения горутин
	wg.Wait()
}

func square(n int) int {
	return n * n
}
