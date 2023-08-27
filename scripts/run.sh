#!/bin/bash

cd "$(dirname "$0")"
cd ..

echo "Stopping Existing Containers..."
docker compose -f zk-single-kafka-single.yml down > /dev/null 2>&1
echo "Starting Containers..."
docker compose -f zk-single-kafka-single.yml up