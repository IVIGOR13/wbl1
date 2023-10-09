package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

func main() {
	arr := []int{7, 3, 8, 34, 8, 7, 73, 4, 6, 8, -7, 65, 9, 0, 3, 2, 1}

	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] < arr[j]
	// })

	quickSort(arr)

	fmt.Println(arr)
}

func quickSort(arr []int) []int {
	return quickSortInterval(arr, 0, len(arr)-1)
}

func quickSortInterval(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortInterval(arr, low, p-1)
		arr = quickSortInterval(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
