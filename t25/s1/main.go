package main

import "time"

func Sleep(d time.Duration) {
	// After выдаст значением из канала по истечении d, что разблокирует выполнение
	// также можно взять реализацию функции After <-NewTimer().C как еще одно решение
	<-time.After(d)
}

func main() {
	Sleep(5 * time.Second)
}
