package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
//
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func AllCharIsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, v := range []rune(strings.ToLower(s)) {
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v] = true
		}
	}
	return true
}

func main() {
	word := "abcd"
	fmt.Printf("%s: %t\n", word, AllCharIsUnique(word))

	word = "aAcd"
	fmt.Printf("%s: %t\n", word, AllCharIsUnique(word))

	word = "abCdefAaf"
	fmt.Printf("%s: %t\n", word, AllCharIsUnique(word))

	word = "aabcd"
	fmt.Printf("%s: %t\n", word, AllCharIsUnique(word))
}
