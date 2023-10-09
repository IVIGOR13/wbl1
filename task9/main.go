package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func main() {
	ch1 := make(chan int64)
	ch2 := make(chan int64)

	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	go func() {
		for val := range ch2 {
			fmt.Println(val)
			wg.Done()
		}
	}()

	go func() {
		for val := range ch1 {
			ch2 <- val * 2
		}
	}()

	for _, v := range arr {
		ch1 <- v
	}

	wg.Wait()

}
