# Mock-Broker-Exchange

## Description
Project mock broker exchange

## Run Frontend
Move to directory `front` and run follow with command.
```
cd frontend
npm install
npm run dev
```
## Run Backend

Move to directory `backend` and run follow with command.

note: change .env file for custom your field

```
cd backend
go mod download
go run main.go
```

## Run full system
Build image `docker` and run with `docker-compose`. You can run follow command

```
cd backend
docker build -t mybroker-api .
cd ../frontend
docker build -t mybroker-front .

docker-compose up 
```
