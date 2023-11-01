package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/carlos-medina/go-data-platform/endpoint"
	"github.com/carlos-medina/go-data-platform/strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap.servers": "localhost:9092", // running locally
		"bootstrap.servers": "broker:29092", // running on docker

		"group.id":          strings.RandStr(rand.Int63(), 10),
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
			input, err := endpoint.DecodeInput(msg.Value)

			if err != nil {
				fmt.Printf("Could not decode message on topic partition: %s\nMessage: %s\nError: %v\n\n", msg.TopicPartition, string(msg.Value), err)
			} else {
				fmt.Printf("Message on topic partition: %s\nDecoded Message: %+v\n\n", msg.TopicPartition, input)
			}
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n\n", err, msg)
		}
	}

	c.Close()
}
