# SkyMP Master API

Another SkyMP Master API implementation written on Go. Supports only PostgreSQL databases. Verification methods are not yet provided.

## Build

### Prerequisites

- [Go](https://go.dev/) for build.
- [Docker](https://www.docker.com/) for build docker image or use local PostgreSQL database.
- [Migrate](https://github.com/golang-migrate/migrate) to load schema in local PostgreSQL database.

### Installing dependencies

```bash
go get -d -v ./
```

### Building app

```bash
go build -o main ./cmd/
```

### Run app without build

```bash
go run ./cmd/
```

### Build docker image

```bash
docker build -t skymp-master-api .
```

### Run docker image

```bash
docker run --name=master -e DB_PASSWORD="DB_PASSWORD_HERE" -e PASSWORD_SALT="PASSWORD_SALT_HERE" -e JWT_SECRET="JWT_SECRET_HERE" -d -p 3000:3000 --rm skymp-master-api
```

## Configuration

App use environment variables and config file.

All environment variables:

- DB_PASSWORD
- PASSWORD_SALT
- JWT_SECRET

Config example (configs/config.json)

```json
{
  "port": 3000,
  "production": false,
  "db_config": {
    "host": "localhost",
    "port": "5432",
    "username": "postgres",
    "db_name": "postgres",
    "ssl_mode": "disable"
  }
}
```

## Development

During development, you can use local database instance in docker via

```bash
docker run --name=postgres-db -e POSTGRES_PASSWORD='YOU_DB_PASSWORD' -d -p 5432:5432 --rm postgres
```

You need to use Migrate to load schema into DB

```bash
migrate -path ./schema -database 'postgres://postgres:12345@localhost/postgres?sslmode=disable' up
```
