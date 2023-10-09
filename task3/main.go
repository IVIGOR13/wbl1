package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	arr := []int64{2, 4, 6, 8, 10}

	fmt.Println(pow2(arr))
}

func pow2(arr []int64) int64 {
	var result int64
	ch := make(chan int64, len(arr))

	for _, val := range arr {
		go func(val int64, ch chan<- int64) {
			ch <- val * val
		}(val, ch)
	}

	for i := 0; i < len(arr); i++ {
		result += <-ch
	}

	return result
}
