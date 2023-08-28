package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} //berfungsi untuk sinkronisasi

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) { //implementasi GoRoutine
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}
