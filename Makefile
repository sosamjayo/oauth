# docker compsoe
up:
	docker compose up -d --remove-orphans
down:
	docker compose down
build:
	docker compose build --no-cache

# Go
lint:
	staticcheck api/...
test:
	go test api/... -v --count=1
