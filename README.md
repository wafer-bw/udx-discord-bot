# udx-discord-bot
A serverless Discord slash command bot powered by Vercel written in Golang using [disgoslash](https://github.com/wafer-bw/disgoslash).

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
# Run all the things you should before you make a commit
make precommit
# Deploy to preview
make preview
# Deploy to production
make deploy
# Sync commands to Discord
make sync
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
# Run all the things you should before you make a commit
go test -coverprofile=cover.out `go list ./... | grep -v ./disgoslash/generatedmocks`
golangci-lint run
gofmt -s -w .
# Deploy to preview
vercel
# Deploy to production
vercel --prod
# Sync commands to Discord
go run sync/sync.go
```

### Developing Slash Commands
1. Make a new folder for your command within [./commands](./commands) using the name of your command.
2. Make a new `.go` script in your new folder for your command.
4. Make sure your new script exports a variable `SlashCommand` like this:
    ```golang
    // SlashCommand instance
    var SlashCommand = slashcommands.New(name, command, do, global, guildIDs)
    ```
    - `name string`: The name of your slash command.
    - `appCommand *ApplicationCommand`: A definition of the [ApplicationCommand object](https://discord.com/developers/docs/interactions/slash-commands#applicationcommand) needed to automatically register the command.
    - `do func`: The function where all your slash command work lives.
    - `global bool`: Whether or not the slash command should be registered globally across all servers your bot has access to.
    - `guildIDs []string`: The guild (server) IDs to register your slash command to.
5. Add your exported `SlashCommand` variable to the list within [./commands/commands.go](./commands/commands.go) like this:
    ```golang
    // SlashCommandMap to be deployed and used live
    var SlashCommandMap = slashcommands.NewMap(
        extrinsicrisk.SlashCommand,
        helloworld.SlashCommand,
        yourcommand.SlashCommand, // <-- Your command goes here.
    )
    ```
6. Open a PR or push to master. Once your changes have been merged/pushed to master they will be automatically deployed to Discord by the [Sync Workflow](./.github/workflows/sync.yml)


## TODOs
* readme
    * table of contents
* `chstrat`
    * handle errors properly and respond in discord
    * use goroutines
        * cancel out and return before 10s vercel time limit
    * write tests
* `sync`
    * move to cmd folder and convert to command
* webhooks
    * Use Discord [webhooks](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks) to notify when a sync happens in the botcommand channel.