# [WIP] go-data-platform

This is a work in progress project. After finished, it will contain the following:

- Ingestor: A kafka worker that consumes from a topic, validates its data and persist it in a database;
- Retriever: A HTTP API that receives a GET request, applies the filter in the query parameter, access the database and returns the data in JSON.

## Running a Kafka broker

Run a Kafka broker on localhost:9092. One way to do this is running [Confluent Platform's Kafka](https://docs.confluent.io/platform/current/platform-quickstart.html) using docker.

## Producing input data

Enter Kafka broker's container:

```bash
docker exec -it broker bash
```

Create a new topic called **input data**:

```bash
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic input-data
```

You can check if the topic was created using the command:

```bash
kafka-topics --list --bootstrap-server localhost:9092
```

In **resources/input-data.txt** in this repo, each line contains a different input. To produce an input event, create a *console producer*, copy and paste one line in the console:

```bash
kafka-console-producer --bootstrap-server localhost:9092 --topic input-data
```

You can check if the event was produced creating a console consumer:

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic input-data --from-beginning
```