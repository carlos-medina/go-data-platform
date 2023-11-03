package main

import (
	"math/rand"

	"github.com/carlos-medina/go-data-platform/strings"

	"github.com/arquivei/foundationkit/errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func MustNewKafkaConsumer() *kafka.Consumer {
	const op = errors.Op("main.MustNewKafkaConsumer")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap.servers": "localhost:9092", // running locally
		"bootstrap.servers": "broker:29092", // running on docker

		"group.id":          strings.RandStr(rand.Int63(), 10),
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(errors.E(op, err))
	}

	err = c.SubscribeTopics([]string{"input-data"}, nil)

	if err != nil {
		panic(errors.E(op, err))
	}

	return c
}
