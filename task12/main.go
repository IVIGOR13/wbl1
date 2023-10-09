package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

var arr []string = []string{"cat", "cat", "dog", "cat", "tree"}

func main() {
	set := make(map[string]bool)

	for _, s := range arr {
		set[s] = true
	}

	for key := range set {
		fmt.Println(key)
	}
}
