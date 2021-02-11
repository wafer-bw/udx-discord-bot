# udx-disco-bot
A serverless discord slash command bot powered by [Vercel](https://vercel.com/) written in [Golang](https://golang.org/)

![tests](https://github.com/wafer-bw/udx-discord-bot/workflows/tests/badge.svg)
![lint](https://github.com/wafer-bw/udx-discord-bot/workflows/lint/badge.svg)

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
# Run tests
go test -coverprofile=cover.out `go list ./... | grep -v ./app/generatedmocks`
# Run verbose tests
go test -v -coverprofile=cover.out `go list ./... | grep -v ./app/generatedmocks`
# Run linting
golangci-lint run
# Run formatting
gofmt -s -w .
# Regenerate mocks
# todo - add `make mock` equivalent
# Run all the things you should before you make a commit
# todo - add `make mock` equivalent
go test -coverprofile=cover.out `go list ./... | grep -v ./app/generatedmocks`
golangci-lint run
gofmt -s -w .
# Deploy to preview
# todo - add `make mock` equivalent
vercel
# Deploy to production
# todo - add `make mock` equivalent
vercel --prod
```

### Command Management

#### List Existing
```sh
go run disgoslash.go list [-v|--verbose]            # Lists global commands
# OR
go run disgoslash.go list <guildID> [-v|--verbose]  # Lists guild commands
```

#### Register New
```sh
go run disgoslash.go create <command-json-path>            # Registers a global command
# OR
go run disgoslash.go create <guildID> <command-json-path>  # Registers a guild command
```

#### Delete Existing
```sh
go run disgoslash.go delete <commandID>            # Delete a global command
# OR
go run disgoslash.go delete <guildID> <commandID>  # Delete a guild command
```

#### Edit
todo

## TODOs
* Readme
    * Document how to use each of the sections and where to code actions for others to use
    * Add badges after the repo is public
        * [Go Report Card](https://goreportcard.com/)
        * [Coveralls](https://coveralls.io/)
        * [CodeQL](https://github.com/wafer-bw/udx-disco-bot/security)
* Code
    * General
        * Check if it's possible to switch from `fmt` to `log`
        * Add scripts that act as an alternative for `make`
    * `commands`
        * Write tests
    * `disgoslash`
        * Write tests
            * [unit test argparsing](https://github.com/docopt/docopt.go/blob/master/examples/unit_test/unit_test.go)
    * `handler`
        * Write tests
    * `client`
        * Handle errors from API responses properly
        * Write tests
        * EditGlobalApplicationCommand
        * EditGuildApplicationCommand
    * `models`
        * Finish Guild Model
    * Decide what to do with `config` within `handler` or `auth` and potentially remove need for `mock`
* Vercel
    * Figure out how to manage env vars
    * Figure out how to manage dev/staging subdomain/branch/deployment
* Log Drain
    * Parse output from stdout and stderr out of log blob
* Bot / Application
    * Give bot an image
* Repo
    * Make public
    * Add branch protection for `master`
* Extract `disgoslash.go` & `app` together into separate repo

## References
* [discordgo](https://github.com/bwmarrin/discordgo) - ed25519 auth
