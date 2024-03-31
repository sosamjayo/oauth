REDIS_ADDR="localhost:6379"
REDIS_PASSWORD="password"

# docker compsoe
up:
	docker compose up -d --remove-orphans
down:
	docker compose down
build:
	docker compose build --no-cache

# Go
run:
	REDIS_ADDR=$(REDIS_ADDR) REDIS_PASSWORD=$(REDIS_PASSWORD) go run ./api/main.go
lint:
	staticcheck ./api/...
test: up
	REDIS_ADDR=$(REDIS_ADDR) REDIS_PASSWORD=$(REDIS_PASSWORD) go test -v ./api/... --count=1 -p=1
check: lint test