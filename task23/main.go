package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.

func DeleteByIndex(arr []int, index int) []int {
	// return slices.Delete(arr, index, index+1)
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	arr = DeleteByIndex(arr, 6)
	fmt.Println(arr)
}
