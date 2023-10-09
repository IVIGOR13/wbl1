package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int64{2, 4, 6, 8, 10}

	pow2(arr)
}

func pow2(arr []int64) {
	wg := sync.WaitGroup{}

	wg.Add(len(arr))

	for _, val := range arr {
		go func(val int64) {
			fmt.Printf("%d^2 = %d\n", val, val*val)
			wg.Done()
		}(val)
	}

	wg.Wait()
}
