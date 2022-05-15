# SkyMP Master API

Red House SkyMP Master API implementation written on Go. Supports only MongoDB databases. Verification methods are not yet provided.

## Build

### Prerequisites

- [Go](https://go.dev/) for build.
- [Docker](https://www.docker.com/) for use local MongoDB database.

### Installing dependencies

```bash
go get -d -v ./
```

### Building app

```bash
go build -o main ./cmd/app/main.go
```

### Run app without build

```bash
go run ./cmd/app/main.go
```

## Configuration

App use environment variables and config file.

All environment variables:

- MONGO_URI - string
- PASSWORD_SALT - string
- JWT_SECRET - string
- PORT - uint16 (0 < port < 65535)
- PRODUCTION - bool

## Development

During development, you can use local database instance in docker via:

```bash
docker run --name mongodb -d -p 27017:27017 mongo
```