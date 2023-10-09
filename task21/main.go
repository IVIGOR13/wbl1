package main

import (
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.

type Client interface {
	Print()
}

func PrintClient(c Client) {
	c.Print()
}

type A struct{}

func (this A) Print() {
	fmt.Println("A")
}

type B struct{}

func (this B) Println() {
	fmt.Println("B")
}

type BAdapter struct {
	B B
}

func (this BAdapter) Print() {
	this.B.Println()
}

func main() {
	PrintClient(A{})
	PrintClient(BAdapter{
		B: B{},
	})
}
