package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 7, 56

	fmt.Printf("Before: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	// a, b = b, a

	fmt.Printf("After: a = %d, b = %d\n", a, b)
}
