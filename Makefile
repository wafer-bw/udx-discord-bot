get:
	go get -t -v -d ./...
.PHONY: get

test:
	go test -coverprofile=cover.out `go list ./... | grep -v ./generatedmocks`
.PHONY: test

testv:
	go test -v -coverprofile=cover.out `go list ./... | grep -v ./generatedmocks`
.PHONY: test

test-ci:
	go test -covermode=count -coverprofile=coverage.out `go list ./... | grep -v ./generatedmocks`
.PHONY: test

lint:
	golangci-lint run
.PHONY: lint

fmt:
	gofmt -s -w .
.PHONY: fmt

tidy:
	go mod tidy
.PHONY: tidy

precommit:
	make get
	make tidy
	make test
	make lint
	make fmt
.PHONY: precommit

preview:
	make precommit
	vercel
.PHONY: preview

deploy:
	make precommit
	vercel --prod
.PHONY: deploy

sync:
	go run sync/sync.go
.PHONY: sync
