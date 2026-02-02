#!/bin/sh

docker-entrypoint.sh postgres &

echo "Waiting for Postgres..."
while ! pg_isready -h localhost -p 5432 > /dev/null 2>1; do
  sleep 1
done

echo "Postgres is up. Starting service"
./transaction-service