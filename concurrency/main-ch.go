package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go process("Order", ch)

	/* step1
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(msg)
	}
	*/

	// step2
	for msg := range ch {
		fmt.Println(msg)
	}
}

func process(item string, ch chan string) {
	defer close(ch)

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		ch <- item
	}
}
