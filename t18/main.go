package main

import (
	"log"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	// блокируем при инкременте
	c.mu.Lock()
	defer wg.Done()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) GetValue() int {
	return c.value
}

func main() {
	wg := &sync.WaitGroup{}
	counter := NewCounter()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Increment(wg)
	}

	// блокируемся и ждем завершения всех горутин
	wg.Wait()

	log.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
