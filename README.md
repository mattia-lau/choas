# CheckList

[x] fetch a BTCUSD pair from any publicly available API, once per minute  
[x] periodically store the exchange rate in a sensible database of your choice  
[x] implement API handlers against the database  
[x] get the last price  
[x] get the price at a given timestamp, come up with a way to serve a price if you don't have price at the requested second  
[x] compute the average price in a time range

# Install Helper Library

```sh
brew install mockery
```

# Build

```sh
docker build -t choas-backend .
docker run -p 8080:8080 choas-backend
```

# Scripts

## Run Test

```sh
make test
```

## Run Server

```sh
make start-server
```

# Run Migration

```sh
go run migrations/migrate_up.go
```

# Generate New mock

```sh
mockery --all
```
