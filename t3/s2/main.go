package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	numbers := []int{2, 4, 6, 8, 10}

	// сюда будем записывать квадраты чисел
	// порядок записи неопределен
	squares := make(chan int, len(numbers))

	for _, n := range numbers {
		wg.Add(1)

		// закидываем на работу горутины
		// каждая горутина записывает квадрат своего числа в squares
		go func(n int) {
			defer wg.Done()
			squares <- square(n)
		}(n)
	}

	// эта горутина нужна, чтобы закрыть канал squares
	// т.к. если его не закрыть то следующий цикл попадет в deadlock
	// (будет ждать следующего значения, но его не будет, т.к. у нас ограниченный набор чисел)
	go func() {
		wg.Wait() // эта строчка "разблокируется", когда квадраты всех чисел посчитаются и запишутся в squares
		close(squares)
	}()

	sum := 0

	// проходимся по квадратам чисел и считаем сумму
	for s := range squares {
		sum += s
	}

	log.Println("SUM", sum)
}

func square(n int) int {
	return n * n
}
