# Greenlight

## Install

### Migrate tool

#### Mac
```
brew install golang-migrate
```

#### Linux
```
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
$ mv migrate.linux-amd64 $GOPATH/bin/migrate
```

## Setup

### Db
Create `greenlight` user
```
CREATE USER greenlight WITH ENCRYPTED PASSWORD 'greenlight';
```
Create `greenlight` database in postgres.
```
CREATE DATABASE greenlight;
```
Grant access to `public schema` if using >=`pg15`
```
\c greenlight
GRANT ALL ON SCHEMA public TO greenlight;
```
Create `citext` extension
```
\c greenlight
create extension citext;
```
Configure environment variable connection:
```
$GREENLIGHT_DB_DSN=<conn_string>
```
Confirm connection:
```
psql $GREENLIGHT_DB_DSN
```

#### Migrations
```
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

# new migration
migrate create -seq -ext=.sql -dir=./migrations create_movies_table
```

## Run
```
go run cmd/api/*.go
```
---


## Calls
### create movie

```
set BODY '{"title":"Moana","year":2016,"runtime":"107 mins","genres":["animation","adventure"]}'
curl -i -d "$BODY" localhost:4000/v1/movies
```

### get movie
```
curl -i localhost:4000/v1/movies/1
```

### update movie
```
set BODY '{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
curl -X PATCH -d "$BODY" localhost:4000/v1/movies/1
```

### create user
```
set BODY '{"name": "Alice Smith", "email": "alice@example.com", "password": "12345678"}'
curl -i -d "$BODY" localhost:4000/v1/users
```


### Call with errors

```
curl -i localhost:4000/foo

curl -i localhost:4000/v1/movies/abc

curl -i -X PUT localhost:4000/v1/healthcheck
```
