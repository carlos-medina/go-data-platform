package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func randStr() string {
	length := 10

	ran_str := make([]byte, 0)

	for i := 0; i < length; i++ {
		ran_str = append(ran_str, byte(65+rand.Intn(25)))
	}

	return string(ran_str)
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap.servers": "localhost:9092", // running locally
		"bootstrap.servers": "broker:29092", // running on docker

		"group.id":          randStr(),
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{"input-data"}, nil)

	if err != nil {
		panic(err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
