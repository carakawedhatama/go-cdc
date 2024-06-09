# go-cdc

This repository demonstrates how to use Change Data Capture (CDC) with Debezium, Kafka, and Go.
Inspired by [this video](https://www.youtube.com/watch?v=5ETH7ENJ-Vs&ab_channel=ProgrammerZamanNow), thanks to [Eko Kurniawan Khannedy](https://www.linkedin.com/in/khannedy/) 🎉.

## Prerequisites

- Docker
- Docker Compose
- Go

## Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/carakawedhatama/go-cdc.git
   cd go-cdc
   ```

2. Start the services using Docker Compose:
   ```bash
   docker-compose up -d
   ```

3. Run the Go application:
   ```bash
   go mod tidy
   go run main.go
   ```

## How it works?
- MariaDB is set up with a sample database and table.
- Debezium captures changes from the MariaDB transaction log and sends them to Kafka.
- The Go-CDC application consumes the changes from the Kafka topic and prints them.

## License
This project is licensed under the MIT License.