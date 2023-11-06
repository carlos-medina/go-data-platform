package main

import (
	"fmt"
	"time"

	"github.com/carlos-medina/go-data-platform/endpoint"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c := MustNewKafkaConsumer()
	mySQLAdapter := MustNewMySQLAdapter()

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			record, err := endpoint.DecodeInput(msg.Value)

			if err != nil {
				fmt.Printf("Could not decode message on partition: %s\nMessage: %s\nError: %v\n\n", msg.TopicPartition, string(msg.Value), err)
			} else {
				fmt.Printf("Message on partition: %s\nDecoded Message: %+v\n\n", msg.TopicPartition, record)

				err := Service(record, mySQLAdapter)
				if err != nil {
					fmt.Printf("Error using service: %v", err)
				} else {
					fmt.Printf("Success on processing records!")
				}
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
