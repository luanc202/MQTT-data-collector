
include .env

run-api:
	go run api/main.go
compose:
	docker compose --env-file .env up --build -d
