// using the builtin sync.Map which is concurrent safe
// alternative: concurrent-map: https://github.com/orcaman/concurrent-map

package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	m := new(sync.Map)

	// конкурентная запись
	go func() {
		m.Store("key1", 42)
	}()

	go func() {
		m.Store("key2", 17)
	}()

	// конкурентное чтение
	go func() {
		value, _ := m.Load("key1")
		log.Println("key1:", value)
	}()

	go func() {
		value, _ := m.Load("key2")
		log.Println("key2:", value)
	}()

	time.Sleep(time.Second)
}
