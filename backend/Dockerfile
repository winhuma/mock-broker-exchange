FROM golang:1.18-alpine

WORKDIR /broker

COPY . /broker
RUN go mod download

EXPOSE 8080

RUN go build -o /mybroker
CMD [ "/mybroker" ]
