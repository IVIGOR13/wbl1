package main

import (
	"fmt"
	"slices"
)

// Реализовать пересечение двух неупорядоченных множеств

var (
	arr1 []int = []int{0, -2, -7, 1, 1, 1, 1, 7, 7, 7, 7, 1234, 5, 4267, 9}
	arr2 []int = []int{1, 9, 4, 7, 2, 1, 5}
)

func main() {
	result := []int{}

	for _, v := range arr1 {
		if slices.Contains(arr2, v) && !slices.Contains(result, v) {
			result = append(result, v)
		}
	}

	for _, v := range arr2 {
		if slices.Contains(arr1, v) && !slices.Contains(result, v) {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
