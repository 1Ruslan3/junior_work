run:
	docker-compose up --build

down:
	docker-compose down -v

swag:
	docker-compose run --rm app swag init

fmt:
	gofmt -w .

test:
	go test ./...

migrate-up:
	docker-compose run --rm migrate

migrate-down:
	docker-compose run --rm migrate -path=/migrations -database=postgres://postgres:postgres@db:5432/subscriptions?sslmode=disable down
