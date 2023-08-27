# Go-kafka-playground

> Reference: https://voskan.host/2023/08/05/golang-apache-kafka/

## Prerequisites

Make sure you have the following software installed:

- Docker / Docker-compose
- Go

## Kafka Setup Using Docker Compose

### 1. Start Docker Compose

First, navigate to the directory containing the Docker Compose file and execute the `scripts/run.sh` script to start the ZooKeeper and Kafka containers.

```bash
./scripts/run.sh
```

### 2. Access the Kafka Container

To access the running Kafka container, execute the following command:

```
docker exec -it kafka1 /bin/bash
```

### 3. Create a Kafka Topic

Once you have accessed the Kafka container, run the following command to create a new topic named "my_topic":

```
kafka-topics --create --topic my_topic --bootstrap-server localhost:19092 --partitions 1 --replication-factor 1
```

- --create: Creates a new topic.
- --topic my_topic: Specifies the name of the topic to create.
- --bootstrap-server localhost:19092: Specifies the Kafka broker address.
- --partitions 1: Specifies the number of partitions for the topic.
- --replication-factor 1: Specifies the replication factor for the topic.

### 4. Verify the Created Topic

If you see the name "my_topic" in the output, the topic has been successfully created.

With these steps, you can create and verify topics in a Kafka environment configured using Docker Compose.

```
kafka-topics --list --bootstrap-server localhost:19092
```

## Execution

### Step 1: Start the Consumer

First, make sure to run the consumer to listen for incoming messages. Navigate to the consumer directory and execute the following command:

```bash
go run consumer/main.go
```

### Step 2: Run the Producer

Once the consumer is up and running, you can proceed to start the producer. Navigate to the producer directory and execute the following command:

```bash
go run producer/main.go
```

By following these steps in sequence, you'll first start the consumer and then run the producer to test your Kafka setup.
