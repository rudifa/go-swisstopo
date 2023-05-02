# Getting information from swisstopo map services

This is a first shot at retrieving information from the swisstopo services.

## Run dev

`go run . [commands] [flags]` _# build an in-memory executable and run it_

## Build and run

`go build .` _# build the executable `go-swisstopo` in the project directory_

`./go-swisstopo [commands] [flags]` _# run it_

## Usage

```
go-swisstopo returns information from the Swiss Federal Office of Topography (swisstopo)

Usage:
  go-swisstopo [command]

Available Commands:
  help        Help about any command
  mapservices Returns a list of map services

Flags:
  -h, --help   help for go-swisstopo

Use "go-swisstopo [command] --help" for more information about a command.
```
