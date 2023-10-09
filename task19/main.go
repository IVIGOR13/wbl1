package main

import (
	"fmt"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func Reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	s := "Hello world! 世界"
	s = Reverse(s)
	fmt.Println(s)

	s = "Ångström"
	s = Reverse(s)
	fmt.Println(s)
}
