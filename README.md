# udx-disco-bot
A serverless discord slash command bot powered by [Vercel](https://vercel.com/) written in [Golang](https://golang.org/)

![tests](https://github.com/wafer-bw/udx-discord-bot/workflows/tests/badge.svg)
![lint](https://github.com/wafer-bw/udx-discord-bot/workflows/lint/badge.svg)

## Getting Started

### Prerequisites
* [Golang](https://golang.org/dl/)
* [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
* [Vercel](https://vercel.com/)
* [Discord](https://discord.com/)
* [Discord Application](https://discord.com/developers/applications)

### Setup
- Clone repo
    ```sh
    git clone git@github.com:wafer-bw/udx-discord-bot.git
    ```
- Get dependencies
    ```sh
    make get
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

### Usage
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
```

### Developing
todo

### Deploying
```sh
# Deploy to preview
make preview
# Deploy to production
make deploy
```

## References
* [discordgo](https://github.com/bwmarrin/discordgo) - ed25519 auth

## TODOs
* Code
    * Make simple scripts for registering commands for now
        - EditGlobalApplicationCommand
        - EditGuildApplicationCommand
    * CLI Tool for slash commands
        * verbose option for list commands
        * handle errors from API responses
        * [unit test argparsing](https://github.com/docopt/docopt.go/blob/master/examples/unit_test/unit_test.go)
    * Add scripts that act as an alternative for `make`
    * Cleanup `handler.go`
        * Modularize
        * Complete tests
    * Finish Guild Model
    * Design command error response flow
    * Check if it's possible to switch from `fmt` to `log`
* Vercel
    *  Figure out how to manage env vars
    * Figure out how to manage dev/staging subdomain/branch/deployment
* Log Drain
    * Parse output from stdout and stderr out of log blob
* Bot / Application
    * Give bot an image
* Repo
    * Make public
    * Add badges requiring the repo be public
        * [Go Report Card](https://goreportcard.com/)
        * [Coveralls](https://coveralls.io/)
        * [CodeQL](https://github.com/wafer-bw/udx-disco-bot/security)
    * Add branch protection for `master`
* Package Repo
    * Extract slash command code to another repo that can act as a library
    * Badges
    * License
