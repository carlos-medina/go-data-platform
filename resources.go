package main

import (
	"database/sql"
	"math/rand"

	"github.com/carlos-medina/go-data-platform/endpoint/gateway"
	"github.com/carlos-medina/go-data-platform/strings"

	"github.com/arquivei/foundationkit/errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-sql-driver/mysql"
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

func MustNewMySQLAdapter() *gateway.MySQLAdapter {
	const op = errors.Op("main.MustNewMySQLAdapter")

	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "172.19.0.2:3306",
		DBName: "go_data_platform",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(errors.E(op, err))
	}

	err = db.Ping()

	if err != nil {
		panic(errors.E(op, err))
	}

	return &gateway.MySQLAdapter{
		DB:    db,
		Table: "records",
	}
}
