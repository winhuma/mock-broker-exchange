# Mock-Broker-Exchange

## Description
Project mock broker exchange

## Run Backend Only

Can move to directory `backend` and run with command.

```
cd backend
go mod download
go run main.go
```

Build image `docker` and run with `docker-compose`. You can run follow command

```
cd backend
docker build -t mybroker .
docker-compose -f backend-compose.yml up 
```
