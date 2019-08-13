DIR:=$(PWD)

.PHONY: build
build:
	go build -o $(DIR)/bin/avsanagrams $(DIR)/cmd/avsanagrams

.PHONY: test
test:
	go test ./...
