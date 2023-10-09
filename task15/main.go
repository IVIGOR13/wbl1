package main

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
	// v := createHugeString(1 << 10)
	// justString = v[:100]

	justString = createHugeString(100)
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
