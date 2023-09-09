check: mock test

mock:
	@echo "Generate Mock..."
	@sh mockgen.sh

test:
	@echo "Run Unit Test..."
	@sh coverage.sh

build:
	@go build cmd/main.go

gorun:
	@go run cmd/main.go

gorun-mq:
	@go run cmd/mq/main.go

docker-start:
	@docker-compose up -d

docker-stop:
	@docker-compose down

run: gorun

run-mq: docker-start gorun-mq