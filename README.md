# goSPFFlare

Go binary for creating/updating MTA-STS records on Cloudflare.

<br>

# PLACEHOLDER
## Generate Cloudflare API Token
1. Visit [https://dash.cloudflare.com/profile/api-tokens](https://dash.cloudflare.com/profile/api-tokens)
2. Create Token
3. "Edit Zone DNS" Template
4. "Zone Resources" Include > Specific Zone > example.com

## Installation via Homebrew (MacOS/Linux - x86_64/arm64)
```
brew install stenstromen/tap/gospfflare
```
## Download and Run Binary
* For **MacOS** and **Linux**: Checkout and download the latest binary from [Releases page](https://github.com/Stenstromen/gospfflare/releases/latest/)
* For **Windows**: Build the binary yourself.

## Build and Run Binary
```
go build
./gospfflare
```

## Example Usage
```
- Usage examples

Go binary for creating and updating SPF record on Cloudflare.

Usage:
  gospfflare [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a new SPF record
  help        Help about any command
  update      Update an existing SPF record

Flags:
  -h, --help   help for gospfflare

Use "gospfflare [command] --help" for more information about a command.
```

<br>

# Random notes

