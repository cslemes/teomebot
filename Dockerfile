FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /app/tmwbot -a -ldflags="-s -w" -installsuffix cgo

FROM scratch AS prod

WORKDIR /app

COPY --from=build /app/tmwbot /

ENTRYPOINT ["/tmwbot"]