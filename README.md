# udx-discord-bot
A serverless Discord slash command bot powered by Vercel written in Golang

![tests](https://github.com/wafer-bw/udx-discord-bot/workflows/tests/badge.svg)
![lint](https://github.com/wafer-bw/udx-discord-bot/workflows/lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/wafer-bw/udx-discord-bot)](https://goreportcard.com/report/github.com/wafer-bw/udx-discord-bot)
![CodeQL](https://github.com/wafer-bw/udx-discord-bot/workflows/CodeQL/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/wafer-bw/udx-discord-bot/badge.svg)](https://coveralls.io/github/wafer-bw/udx-discord-bot)

## Getting Started

### Prerequisites
#### Primary
* [Golang](https://golang.org/dl/)
* [Vercel](https://vercel.com/)
* [Discord](https://discord.com/)
* [Discord Application](https://discord.com/developers/applications)

#### Dev
* [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
* [mockery](https://github.com/vektra/mockery)

### Setup
- Clone repo
    ```sh
    git clone git@github.com:wafer-bw/udx-discord-bot.git
    ```
- Get dependencies
    ```sh
    go get -t -v -d ./...
    ```
- Make `.env` file from sample
    ```sh
    cp .env.sample .env
    ```
- Add application & bot secrets to `.env` file which can be found in the "General Information" and "Bot" sections of your Discord Application page.
    - `CLIENT_ID`
    - `CLIENT_SECRET`
    - `PUBLIC_KEY`
    - `TOKEN`

### Usage (POSIX)
```sh
# Get Dependencies
make get
# Tidy go.mod
make tidy
# Run tests
make test
# Run verbose tests
make testv
# Run linting
make lint
# Run formatting
make fmt
# Regenerate mocks
make mocks
# Run all the things you should before you make a commit
make precommit
# Deploy to preview
make preview
# Deploy to production
make deploy
```

### Usage (Windows) (WIP)
```sh
# Get Dependencies
go get -t -v -d ./...
# Tidy go.mod
go mod tidy
# Run tests
go test -coverprofile=cover.out `go list ./... | grep -v ./disgoslash/generatedmocks`
# Run verbose tests
go test -v -coverprofile=cover.out `go list ./... | grep -v ./disgoslash/generatedmocks`
# Run linting
golangci-lint run
# Run formatting
gofmt -s -w .
# Regenerate mocks
# todo - add `make mock` equivalent
# Run all the things you should before you make a commit
# todo - add `make mock` equivalent
go test -coverprofile=cover.out `go list ./... | grep -v ./disgoslash/generatedmocks`
golangci-lint run
gofmt -s -w .
# Deploy to preview
# todo - add `make mock` equivalent
vercel
# Deploy to production
# todo - add `make mock` equivalent
vercel --prod
```

## TODOs
* Code
    * General
        * Check if it's possible to switch from `fmt` to `log` while using vercel
        * Add scripts that act as an alternative for `make`
    * `client`
        * EditGlobalApplicationCommand
        * EditGuildApplicationCommand
    * `models`
        * Finish Guild Model
    * `exporter`
        * Recreate as package
* Bot / Application
    * Give bot an image
* Extract `disgoslash` into a separate repo
* Resync Workflow

## References
* [discordgo](https://github.com/bwmarrin/discordgo) - ed25519 auth
