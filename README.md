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
