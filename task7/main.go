package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	data := make(map[string]int)
	var mu sync.Mutex

	numGoroutines := 5

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// захват мьютекса
			mu.Lock()
			defer mu.Unlock()

			// запись в мапу
			key := fmt.Sprintf("key_%d", id)
			data[key] = id

			fmt.Printf("[%d]: %s -> %d\n", id, key, id)

			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(data)
}
