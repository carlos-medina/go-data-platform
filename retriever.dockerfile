FROM golang:1.21.0-alpine3.18 as builder
RUN apk add alpine-sdk
WORKDIR /go/app
COPY . /go/app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -C cmd/retriever -o retriever

FROM alpine:latest as runner
WORKDIR /go/app
COPY --from=builder /go/app/cmd/retriever/retriever .
ENTRYPOINT ./retriever