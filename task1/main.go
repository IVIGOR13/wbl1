package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

type Human struct {
	Name string
}

func (h Human) getName() string {
	return h.Name
}

type Action struct {
	Human
}

func main() {
	a := Action{Human: Human{Name: "John Spenser"}}
	fmt.Println(a.getName())
}
