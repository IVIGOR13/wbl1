package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// select
// 	 	context
// 		time.After
// close readed ch
// signal ch

func ex1(ctx context.Context) {
	fmt.Println("start")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println("close")
	}
}

func ex2(done chan struct{}) {
	fmt.Println("start")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	case <-done:
		fmt.Println("close")
	}
}

func ex3(done chan struct{}) {
	fmt.Println("start")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	case <-done:
		fmt.Println("close")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go ex1(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	done := make(chan struct{})
	go ex2(done)
	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)

	done = make(chan struct{})
	go ex3(done)
	time.Sleep(1 * time.Second)
	done <- struct{}{}
	time.Sleep(1 * time.Second)
}
