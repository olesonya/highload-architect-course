FROM golang:1.22.3-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd

FROM alpine:3.19.1

WORKDIR /srv

COPY scripts scripts

COPY --from=builder /usr/local/bin/app /srv/app

ENTRYPOINT ["/srv/app"]
