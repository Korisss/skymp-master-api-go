# skymp-master-api-go

run docker

```bash
docker run --name=skymp-master-api -e POSTGRES_PASSWORD='12345' -d -p 5432:5432 --rm postgres
```

run migrate up

```bash
migrate -path ./schema -database 'postgres://postgres:12345@localhost/postgres?sslmode=disable' up
```

run migrate down

```bash
migrate -path ./schema -database 'postgres://postgres:12345@localhost/postgres?sslmode=disable' down
```

install gin and sqlx

DB_PASSWORD, JWT_SECRET and PASSWORD_SALT in .env

build via docker

```bash
docker build -t skymp-master-api .
```

run via docker

```bash
docker run --name=master -e DB_PASSWORD="DB_PASSWORD_HERE" -e PASSWORD_SALT="PASSWORD_SALT_HERE" -e JWT_SECRET="JWT_SECRET_HERE" -d -p 3000:3000 --rm skymp-master-api
```

config.json file in ./configs/

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
