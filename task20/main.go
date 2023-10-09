package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func Reverse(s []string) []string {
	res := []string(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func reverseWords(input string) string {
	words := strings.Fields(input) // split string by spaces
	words = Reverse(words)         // reverse array of words
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
