#!/bin/sh

# 1. Start PostgreSQL in the background
docker-entrypoint.sh postgres &

# 2. Wait for Postgres to be ready
echo "Waiting for Postgres..."
while ! pg_isready -h localhost -p 5432 > /dev/null 2>1; do
  sleep 1
done

# 3. Start your Go application
echo "Postgres is up. Starting service"
./transaction-service