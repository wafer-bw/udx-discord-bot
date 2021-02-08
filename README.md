# udx-disco-bot
A serverless discord bot powered by [Vercel](https://vercel.com/) written in [Golang](https://golang.org/)

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
todo

### Usage
```sh
# Get dependencies
make get
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

### Deploying
todo

## References
* [discordgo](https://github.com/bwmarrin/discordgo) - ed25519 auth

## TODOs
* Bot / Application
    * Give bot an image
* Repo
    * Make public
    * Add badges requiring the repo be public
        * [Go Report Card](https://goreportcard.com/)
        * [Coveralls](https://coveralls.io/)
        * [CodeQL](https://github.com/wafer-bw/udx-disco-bot/security)
    * Add branch protection for `master`
    * License
* Code
    * Models
        * Guild struct
        * GuildMember struct
        * Embeds struct
        * AllowMentions struct
    * Make simple scripts for registering commands for now
    * Add scripts that act as an alternative for `make`
    * Cleanup `handler.go`
        * Modularize
        * Complete tests
    * Revisit & redesign command error response flow
    * Design command register flow
    * Extract slash command code to another repo that can act as a library
