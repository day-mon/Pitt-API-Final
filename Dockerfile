
FROM golang:1.18-alpine AS build

RUN apk update && apk add ca-certificates && apk add tzdata

ARG VERSION=dev

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-X 'main.version=${VERSION}'-w -s" -o api main.go

FROM alpine

COPY --from=build /build/api /opt/api/api

ENTRYPOINT ["/opt/api/api"]
