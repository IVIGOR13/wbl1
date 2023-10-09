package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func main() {
	a := big.NewFloat(58275939729435)
	b := big.NewFloat(2779235235985)

	c := big.NewFloat(0)

	// умножение
	c = c.Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, c)

	// деление
	c = c.Quo(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, c)

	// вычитание
	c = c.Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, c)

	// сложение
	c = c.Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, c)
}
