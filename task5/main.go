package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться

const (
	N = 1
)

func main() {
	ch := make(chan int64)

	go func() {
		for val := range ch {
			fmt.Printf("%d\n", val)
		}
	}()

	timeout := time.After(N * time.Second)

	for {
		select {
		case <-timeout:
			close(ch)
			fmt.Println("Timeout")
			return
		default:
			ch <- rand.Int63()
		}
	}
}
