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

### Developing Slash Commands
1. Make a new folder for your command within [./commands](./commands) using the name of your command.
2. Make a new `.go` script in your new folder for your command.
3. Use the existing commands as referance for what your script will need.
4. Make sure your new script exports a variable `SlashCommand` like this:
    ```golang
    // SlashCommand - the slash command instance
    var SlashCommand = slashcommands.New(name, command, hello, global, guildIDs)
    ```
5. Add your exported `SlashCommand` variable to list within [./commands](./commands/commands.go) like this:
    ```golang
    // SlashCommandMap for the application
    var SlashCommandMap = slashcommands.NewMap(
        extrinsicrisk.SlashCommand,
        helloworld.SlashCommand,
        yourcommand.SlashCommand,
    )
    ```
6. Open a PR or push to master. Once your changes have been merged/pushed to master they will be automatically deployed to discord by the [Sync Workflow](./.github/workflows/sync.yml)

## TODOs
* Code
    * General
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

## References
* [discordgo](https://github.com/bwmarrin/discordgo) - ed25519 auth
