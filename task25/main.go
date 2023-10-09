package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func Sleep(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	fmt.Println("start")
	Sleep(1500 * time.Millisecond)
	fmt.Println("end")
}
