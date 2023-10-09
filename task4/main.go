package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

const (
	N = 2
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int64)

	for i := 0; i < N; i++ {
		go func(i int) {
			for val := range ch {
				fmt.Printf("[%d] %d\n", i, val)
			}
		}(i)
	}

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt
		fmt.Println("Ctrl-C was pressed")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- rand.Int63()
		}
	}

}
