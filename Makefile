lint:
	golangci-lint --config=$(CURDIR)/golangci.yaml run ./...