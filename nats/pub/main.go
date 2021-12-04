package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	for i := 1; i <= 100; i++ {
		s := fmt.Sprintf("Message %v, data: %v\n", i, random.Intn(10000))
		nc.Publish("events.new", []byte(s))
	}
}
