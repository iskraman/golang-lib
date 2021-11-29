package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			ch1 <- "Order processed"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			ch2 <- "Refund processed"
		}
	}()

	/* step1
	for {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
	*/

	// step2
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
