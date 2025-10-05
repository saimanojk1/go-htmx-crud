#!/bin/bash

if docker ps -a --filter "name=basic-crud" --format "{{.Names}}" | grep -q "basic-crud"; then
  if docker ps --filter "name=basic-crud" --format "{{.Names}}" | grep -q "basic-crud"; then
    echo "Container is already running"
  else
    echo "Starting existing container"
    docker start basic-crud
  fi
else
  echo "Creating and starting new container"
  # docker run -d --name basic-crud -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
  docker run -d --name basic-crud -e POSTGRES_HOST_AUTH_METHOD=trust -p 15432:15432 postgres -p 15432

fi
sleep 10
pg_isready -h localhost -p 15432 -U postgres
psql -h localhost -p 15432 -U postgres -d postgres -f schema.sql
