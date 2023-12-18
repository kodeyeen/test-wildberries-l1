package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// завершение с помощью контекста с таймаутом
	// по сути тоже, что и в предыдущем решении,
	// только сигнал посылается через определенное время, а не когда мы захотим
	// также можно использован context WithCancel, чтобы посылать сигнал вызовом cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go routine(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}

func routine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // по истечении таймаута придет сигнал
			log.Println("Finished")
			return
		default:
			log.Println("Doing something...")
		}
	}
}
