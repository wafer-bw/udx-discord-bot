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

### Slash Commands

#### Developing Commands
- Commands are located within [./udx-slash-commands/commands](./udx-slash-commands/commands).
- Commands must be a `.go` file that resides in it's own folder within the `./udx-slash-commands/commands` directory
    - Ex: `helloworld/helloworld.go`
- When a user executes a slash command it will be received as an object of type `InteractionRequest`
    - Found within [./app/models/interaction.go](./app/models/interaction.go)
- When responding to a slash command we respond with an object of type `InteractionResponse`
    - Found within [./app/models/interaction.go](./app/models/interaction.go)
- Commands must be registered to Discord the data needed to register is defined in an object of type `ApplicationCommand`
    - Found within [./app/models/interaction.go](./app/models/interaction.go)
- A Command `.go` file must export a variable named `SlashCommand` which is created using the method `commands.NewSlashCommand()` which requires a `name (string)`, `command (ApplicationCommand)`, and a function with the signature `func someName(request *models.InteractionRequest) (*models.InteractionResponse, error)` which is the function where you put all your code that does the work you want.
- There are two examples that exist already:
    - [helloworld](./udx-slash-commands/commands/helloworld/helloworld.go)
    - [extrinsicrisk](./udx-slash-commands/commands/extrinsicrisk/extrinsicrisk.go)
- The `SlashCommand` variable must be added to the array of slash commands inside [./udx-slash-commands/commands/commands.go](./udx-slash-commands/commands/commands.go)
- The above code changes must be merged into the `master` branch and pushed to GitHub.
- You must then [export the command](#export-commands)
- You must then [register the command](#register-an-exported-command)

#### Export Commands
After creating a command we must export it to a JSON file which we later use to register the command in Discord.
To export all commands within the `./udx-slash-commands/commands/` directory `cd` into `udx-slash-commands` and run `go run export.go` which will export all commands into JSON files within `./udx-slash-commands/raw`.

#### Register an Exported Command
Commands must be registerd with Discord in order for Discord to start supporting the slash command.
- To register an exported command globally on your application use
    ```sh
    go run disgoslash.go register <pathToExportedCommand>
    ```
    - Global commands take up to an hour for changes to be made live by Discord.
- To register an exported command to a specific guild (server) use
    ```sh
    go run disgoslash.go register <serverID> <pathToExportedCommand>
    # Example:
    # go run disgoslash.go register 12345 udx-slash-commands/raw/helloworld.json
    ```

#### Listing Registered Commands
- Global
    ```sh
    go run disgoslash.go list
    ```
    - Guild / Server commands will not appear here, they are separate from global commands.
- Guild / Server
    ```sh
    go run disgoslash.go list <serverID>
    # Example:
    # go run disgoslash.go list 12345
    ```

#### Unregister a Command
- Global
    ```sh
    go run disgoslash.go unregister <commandID> # Command ID found by listing the registered commands
    ```
    - Guild / Server commands will not appear here, they are separate from global commands.
- Guild / Server
    ```sh
    go run disgoslash.go unregister <serverID> <commandID>
    # Example:
    # go run disgoslash.go unregister 67890 12345
    ```

## TODOs
* Code
    * Get rid of exporter and add a CLI to register that uses exported slash command array
    * General
        * Check if it's possible to switch from `fmt` to `log`
        * Add scripts that act as an alternative for `make`
    * `disgoslash`
        * Write tests
            * [unit test argparsing](https://github.com/docopt/docopt.go/blob/master/examples/unit_test/unit_test.go)
    * `handler`
        * Write tests
    * `client`
        * Handle errors from API responses properly
        * EditGlobalApplicationCommand
        * EditGuildApplicationCommand
        * Write tests
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
* Extract `disgoslash.go` & `app` together into separate repo

## References
* [discordgo](https://github.com/bwmarrin/discordgo) - ed25519 auth
