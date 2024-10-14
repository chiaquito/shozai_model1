.PHONY: run-local
run:
	go run ./main.go


.PHONY: build
build:
	docker compose build

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down -v

