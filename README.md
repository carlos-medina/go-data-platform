# [WIP] go-data-platform

This is a work in progress project. After finished, it will contain the following:

- Ingestor: A kafka worker that consumes from a topic, validates its data and persist it in a database;
- Retriever: A HTTP API that receives a GET request, applies the filter in the query parameter, access the database and returns the data in JSON.