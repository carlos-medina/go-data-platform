# [WIP] go-data-platform

This is a work in progress project. After finished, it will contain the following:

- Ingestor: A kafka worker that consumes from a topic, validates its data and persist it in a database;
- Retriever: A HTTP API that receives a GET request, applies the filter in the query parameter, access the database and returns the data in JSON.

## Running a Kafka broker

Run a Kafka broker on localhost:9092 using docker. The one present in this repo was taken from [Confluent Platform's Kafka](https://docs.confluent.io/platform/current/platform-quickstart.html).

```bash
docker compose up -d kafka-broker
```

## Producing input data

Enter Kafka broker's container:

```bash
docker exec -it kafka-broker bash
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

## Running Ingestor

To run the docker image, the key "bootstrap.servers" on kafka.ConfigMap must have the value "broker:29092" if you are running [Confluent Platform's Kafka](https://docs.confluent.io/platform/current/platform-quickstart.html); more on **KAFKA_LISTENERS** AND **KAFKA__ADVERTISED_LISTENERS** [here](https://stackoverflow.com/questions/61990336/kafka-consumer-failed-to-start-connection-refused-connect2-for-127-0-0-1) and [here](https://rmoff.net/2018/08/02/kafka-listeners-explained/):

```bash
docker compose up ingestor
```