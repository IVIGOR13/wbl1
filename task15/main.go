package main

import (
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.
//
// 	var justString string
// 	func someFunc() {
// 		v := createHugeString(1 << 10)
// 		justString = v[:100]
// 	}
//
// 	func main() {
// 		someFunc()
// 	}

var justString string

func someFunc() {
	str := createHugeString(1 << 10)

	strLen := 100

	builder := strings.Builder{}
	builder.Grow(strLen)
	for _, v := range []byte(str) {
		builder.WriteByte(v)
	}
	justString = builder.String()

	// justBytes := make([]byte, strLen)
	// for i, v := range []byte(str) {
	// 	justBytes[i] = v
	// }
	// justString = string(justBytes)
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = '#'
	}
	return string(bytes)
}
