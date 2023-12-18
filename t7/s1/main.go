package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
	return value, ok
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func main() {
	sm := NewSafeMap()

	// конкурентная запись
	go func() {
		sm.Set("key1", 42)
	}()

	go func() {
		sm.Set("key2", 17)
	}()

	// конкурентное чтение
	go func() {
		value, _ := sm.Get("key1")
		fmt.Println("key1:", value)
	}()

	go func() {
		value, _ := sm.Get("key2")
		fmt.Println("key2:", value)
	}()

	time.Sleep(time.Second)
}
