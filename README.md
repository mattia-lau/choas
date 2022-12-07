# CheckList

[x] fetch a BTCUSD pair from any publicly available API, once per minute  
[x] periodically store the exchange rate in a sensible database of your choice  
[x] implement API handlers against the database  
[x] get the last price  
[x] get the price at a given timestamp, come up with a way to serve a price if you don't have price at the requested second  
[x] compute the average price in a time range

# Install Mockery

```sh
brew install mockery
```

# Build

## Docker
```sh
docker build -t choas-backend .
docker run -p 8080:8080 -v "$(pwd)/chaos.db:/app/choas.db" choas 
```

## Docker Compose
```sh
docker-compose up -d
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

## Run Migration

```sh
go run migrations/migrate_up.go
```

## Generate New mock

```sh
mockery --all
```

# API Document

## Get the last price

```
GET /last/ticker/<symbol>
```

| Slug   | Type   | Description              |
| ------ | ------ | ------------------------ |
| symbol | string | **Required**. eg: BTCUSD |

### Response Example

```json
{
  "ID": 10,
  "Symbol": "BTCUSD",
  "Price": 16988.4346,
  "Date": "2022-12-07T14:12:00.564733+08:00"
}
```

### Status Code

| Status Code | Description    |
| ----------- | -------------- |
| 200         | OK             |
| 400         | Incorrect Slug |

## Get the price at a given timestamp

```
GET /<date>/ticker/<symbol>
```

| Slug   | Type   | Description                                                 |
| ------ | ------ | ----------------------------------------------------------- |
| date   | string | **Required**. In RFC3339 standard, eg: 2022-12-06T00:03:00Z |
| symbol | string | **Required**. eg: BTCUSD                                    |

### Response Example

```json
{
  "ID": 8,
  "Symbol": "BTCUSD",
  "Price": 16994.5,
  "Date": "2022-12-06T00:03:00Z"
}
```

### Status Code

| Status Code | Description    |
| ----------- | -------------- |
| 200         | OK             |
| 400         | Incorrect Slug |

## Compute the average price in a time range

```
GET /average/ticker/<symbol>?start=<start>&end=<end>
```

| Parameter | Type   | Description                                                 |
| --------- | ------ | ----------------------------------------------------------- |
| start     | string | **Required**. In RFC3339 standard, eg: 2022-12-06T00:03:00Z |
| end       | string | **Required**. In RFC3339 standard, eg: 2022-12-06T00:03:00Z |

| Slug   | Type   | Description              |
| ------ | ------ | ------------------------ |
| symbol | string | **Required**. eg: BTCUSD |

### Response Example

```json
{ "average": 17001.909992307694 }
```

### Status Code

| Status Code | Description                |
| ----------- | -------------------------- |
| 200         | OK                         |
| 400         | Incorrect Slug / parameter |

## Example

```sh
curl http://localhost:8080/average/ticker/BTCUSD?start=2022-01-01T00:00:00Z&end=2022-12-31T23:59:59Z
curl http://localhost:8080/2022-12-06T00:03:00Z/ticker/BTCUSD
curl http://localhost:8080/last/ticker/BTCUSD
```
