package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		process("Order")
		wg.Done()
	}()

	wg.Wait()
}

func process(item string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processed", i, item)
	}
}
