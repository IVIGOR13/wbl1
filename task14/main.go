package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func getType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "unknown"
	}
}

func getTypeReflect(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func main() {
	var a interface{} = "Hello"
	var b interface{} = 42
	var c interface{} = true
	var d interface{} = make(chan int)

	fmt.Printf("a: %s\n", getType(a))
	fmt.Printf("b: %s\n", getType(b))
	fmt.Printf("c: %s\n", getType(c))
	fmt.Printf("d: %s\n", getType(d))

	fmt.Println()

	fmt.Printf("a: %s\n", getTypeReflect(a))
	fmt.Printf("b: %s\n", getTypeReflect(b))
	fmt.Printf("c: %s\n", getTypeReflect(c))
	fmt.Printf("d: %s\n", getTypeReflect(d))
}
