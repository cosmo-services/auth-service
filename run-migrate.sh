#!/bin/sh
echo "Debug: PG_USER=$PG_USER"
echo "Debug: PG_PASS=$PG_PASS"
echo "Debug: PG_PORT=$PG_PORT"
echo "Debug: PG_NAME=$PG_NAME"

migrate -path /migrations -database "postgres://${PG_USER}:${PG_PASS}@postgres:${PG_PORT}/${PG_NAME}?sslmode=disable" up