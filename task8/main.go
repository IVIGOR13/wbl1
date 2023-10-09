package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0

func setBit(val *int64, bit, pos int64) {
	if bit == 0 {
		*val = *val & (*val ^ (1 << pos))
	} else if bit == 1 {
		*val |= (1 << pos)
	}
}

func main() {
	var val int64 = 673
	fmt.Printf("%b\n", val)
	setBit(&val, 0, 5)
	fmt.Printf("%b\n", val)
}
