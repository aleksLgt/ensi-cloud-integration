lint:
	golangci-lint --config=$(CURDIR)/golangci.yaml run ./...

build:
	docker-compose build

run:
	docker-compose up --force-recreate --build