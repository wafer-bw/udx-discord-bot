get:
	go get -t -v -d ./...
.PHONY: get

mocks:
	rm -rf app/generatedmocks/*
	mockery --all --output="app/generatedmocks" --keeptree
.PHONY: mocks

test:
	go test -coverprofile=cover.out `go list ./... | grep -v ./app/generatedmocks`
.PHONY: test

testv:
	go test -v -coverprofile=cover.out `go list ./... | grep -v ./app/generatedmocks`
.PHONY: test

test-ci:
	go test -covermode=count -coverprofile=coverage.out `go list ./... | grep -v ./app/generatedmocks`
.PHONY: test

lint:
	golangci-lint run
.PHONY: lint

fmt:
	gofmt -s -w .
.PHONY: fmt

precommit:
	make mocks
	make test
	make lint
	make fmt
.PHONY: precommit
