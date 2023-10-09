package main

import (
	"fmt"
	"slices"
)

// Реализовать бинарный поиск встроенными методами языка

func main() {
	arr := []int{-7, 0, 1, 2, 3, 3, 4, 6, 7, 7, 8, 8, 8, 9, 34, 65, 73}

	val := 9
	index, result := slices.BinarySearch[[]int, int](arr, val)
	fmt.Printf("find %d: result: index = %d finded = %t\n", val, index, result)

	// index, result = binarySearch(arr, val)
	// fmt.Printf("find %d: result: index = %d finded = %t\n", val, index, result)
}

// func binarySearch(arr []int, num int) (int, bool) {

// 	l, r := 0, len(arr)-1

// 	for l <= r {
// 		mid := (l + r) / 2

// 		if arr[mid] < num {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}

// 	if l == len(arr) || arr[l] != num {
// 		return l, false
// 	}

// 	return r + 1, true
// }
