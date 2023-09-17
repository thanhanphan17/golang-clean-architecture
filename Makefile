init_db:
	@docker run \
		--name postgres-db \
		--rm -e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=0000 \
		-e POSTGRES_DB=db-go-clean-architecture\
		-v ./postgres-data:/var/lib/postgresql/data \
		-p 5432:5432 -it \
		-d postgres:latest

rm_db:
	@docker rm -f postgres-db

compose_up:
	@docker compose up

compose_down:
	@docker compose down -v
	@docker rmi -f go-clean-architecture:1.0

local:
	@go run cmd/main.go -config ./config/env -env=local -upgrade=${m}

