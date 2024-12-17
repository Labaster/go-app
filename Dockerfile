FROM golang:1.23.3 AS build

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY .env .

COPY dbConf dbConf
COPY routeActions routeActions
COPY structures structures

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM golang:1.23.3-alpine

WORKDIR /app

COPY --from=build /docker-gs-ping /docker-gs-ping
COPY --from=build /app/.env .env

EXPOSE 4400

# Run
CMD ["/docker-gs-ping"]