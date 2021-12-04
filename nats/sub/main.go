package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	/* cast 1
	sub, _ := nc.Subscribe("events.local", func(msg *nats.Msg) {
		fmt.Printf("Message recv on subject: %v, data: %v\n", msg.Subject, string(msg.Data))
	})

	time.Sleep(10 * time.Second)
	*/

	sub, err := nc.SubscribeSync("events.*")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if msg, err := sub.NextMsg(10 * time.Second); err == nil {
			fmt.Printf("Message recv on subject: %v, data: %v\n", msg.Subject, string(msg.Data))
		} else {
			break
		}
	}

	sub.Unsubscribe()
}
